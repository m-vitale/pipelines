import tensorflow as tf
from tensorflow.python.saved_model import signature_constants
from tensorflow.python.saved_model import tag_constants
from nets import nets_factory
from tensorflow.python.platform import gfile
import argparse
import validators
import os
import requests
import tarfile
from subprocess import call, Popen, PIPE
import shutil
import glob
import re
from  tensorflow.python.tools.freeze_graph import freeze_graph
from  tensorflow.python.tools.saved_model_cli import _show_all
from urllib.parse import urlparse
from shutil import copyfile

def delete_gcs_folder(path):
    parsed_path = urlparse(path)
    bucket_name = parsed_path.netloc
    model_directory = parsed_path.path[1:]
    gs_client = storage.Client()
    bucket = gs_client.get_bucket(bucket_name)
    blobs = bucket.list_blobs(prefix=model_directory)
    for blob in blobs:
        blob.delete()

def upload_to_gcs(src, dst):
    parsed_path = urlparse(dst)
    bucket_name = parsed_path.netloc
    file_path = parsed_path.path[1:]
    gs_client = storage.Client()
    bucket = gs_client.get_bucket(bucket_name)
    blob = bucket.blob(file_path)
    blob.upload_from_filename(src)

def main():
    parser = argparse.ArgumentParser(
        description='Slim model generator')
    parser.add_argument('--model_name', type=str,
                        help='')
    parser.add_argument('--export_dir', type=str, default="/tmp/export_dir",
                        help='GCS or local path to save graph files')
    parser.add_argument('--saved_model_dir', type=str,
                        help='GCS or local path to save the generated model')
    parser.add_argument('--batch_size', type=str, default=1,
                        help='batch size to be used in the exported model')
    parser.add_argument('--checkpoint_url', type=str,
                        help='URL to the pretrained compressed checkpoint')
    parser.add_argument('--num_classes', type=int, default=1000,
                        help='number of model classes')
    args = parser.parse_args()

    MODEL = args.model_name
    URL = args.checkpoint_url
    if not validators.url(args.checkpoint_url):
        print('use a valid URL parameter')
        exit(1)
    TMP_DIR = "/tmp/slim_tmp"
    NUM_CLASSES = args.num_classes
    BATCH_SIZE = args.batch_size
    MODEL_FILE_NAME = URL.rsplit('/', 1)[-1]
    EXPORT_DIR = args.export_dir
    SAVED_MODEL_DIR = args.saved_model_dir

    if not os.path.exists(TMP_DIR):
        os.makedirs(TMP_DIR)

    if not os.path.exists(TMP_DIR + '/' + MODEL_FILE_NAME):
        print("Downloading and decompressing the model checkpoint...")
        response = requests.get(URL, stream=True)
        with open(TMP_DIR + '/' + MODEL_FILE_NAME,'wb') as output:
            output.write(response.content)
            tar = tarfile.open(TMP_DIR + '/' + MODEL_FILE_NAME)
            tar.extractall(path=TMP_DIR)
            tar.close()
            print("Model checkpoint downloaded and decompressed to:",TMP_DIR )
    else:
        print("Reusing existing model file ",TMP_DIR + '/' + MODEL_FILE_NAME)

    checkpoint = glob.glob(TMP_DIR + '/*.ckpt*')
    print("checkpoint",checkpoint)
    if len(checkpoint) > 0:
        m = re.match("([\S]*.ckpt)",checkpoint[-1])
        print("checkpoint match", m)
        checkpoint = m[0]
        print(checkpoint)
    else:
        print("checkpoint file not detected in " + URL)
        exit(1)


    print("Saving graph def file")
    with tf.Graph().as_default() as graph:

        network_fn = nets_factory.get_network_fn(MODEL,
                                                 num_classes=NUM_CLASSES,
                                                 is_training=False)
        image_size = network_fn.default_image_size
        if BATCH_SIZE == "None" or  BATCH_SIZE == "-1":
            batchsize = None
        else:
            batchsize = BATCH_SIZE
        placeholder = tf.placeholder(name='input', dtype=tf.float32,
                                     shape=[batchsize, image_size,
                                            image_size, 3])
        network_fn(placeholder)
        graph_def = graph.as_graph_def()

        with gfile.GFile(os.path.join(TMP_DIR,MODEL + '_graph.pb'), 'wb') as f:
            f.write(graph_def.SerializeToString())
    if urlparse(EXPORT_DIR).scheme == 'gs':
        upload_to_gcs(os.path.join(TMP_DIR,MODEL + '_graph.pb'), os.path.join(EXPORT_DIR,MODEL + '_graph.pb'))
    elif urlparse(EXPORT_DIR).scheme == '':
        if not os.path.exists(EXPORT_DIR):
            os.makedirs(EXPORT_DIR)
        copyfile(os.path.join(TMP_DIR,MODEL + '_graph.pb'), os.path.join(EXPORT_DIR,MODEL + '_graph.pb'))
    else:
        print("Invalid format of model export path")
    print("Graph file saved to ", os.path.join(EXPORT_DIR,MODEL + '_graph.pb'))


    print("Analysing graph")
    p = Popen("./summarize_graph --in_graph=" + os.path.join(TMP_DIR,MODEL + '_graph.pb') + " --print_structure=false",
              shell=True, stdout=PIPE, stderr=PIPE)
    summary, err = p.communicate()
    inputs = []
    outputs = []
    for line in summary.split(b'\n'):
        if re.match("Found [\d]* possible inputs", line.decode()) is not None:
            print("in",line)
            m = re.findall('name=[\S]*,', line.decode())
            for match in m:
                print("match",match)
                input = match[5:-1]
                inputs.append(input)
            print("inputs",inputs)

        if re.match("Found [\d]* possible outputs", line.decode()) is not None:
            print("out",line)
            m = re.findall('name=[\S]*,', line.decode())
            for match in m:
                print("match",match)
                output = match[5:-1]
                outputs.append(output)
            print("outputs",outputs)


    output_node_names = ",".join(outputs)
    print("Creating freezed graph based on pretrained checkpoint")
    freeze_graph(input_graph=os.path.join(TMP_DIR,MODEL + '_graph.pb'), \
                 input_checkpoint=checkpoint, \
                 input_binary=True, \
                 clear_devices=True, \
                 input_saver='', \
                 output_node_names=output_node_names, \
                 restore_op_name="save/restore_all", \
                 filename_tensor_name="save/Const:0", \
                 output_graph=os.path.join(EXPORT_DIR,'frozen_graph_'+MODEL+'.pb'), \
                 initializer_nodes="")
    if urlparse(SAVED_MODEL_DIR).scheme == '' and os.path.exists(SAVED_MODEL_DIR):
        shutil.rmtree(saved_model_dir)

    builder = tf.saved_model.builder.SavedModelBuilder(SAVED_MODEL_DIR)

    with tf.gfile.GFile(os.path.join(EXPORT_DIR,'frozen_graph_'+MODEL+'.pb'), "rb") as f:
        graph_def = tf.GraphDef()
        graph_def.ParseFromString(f.read())

    sigs = {}

    with tf.Session(graph=tf.Graph()) as sess:
        tf.import_graph_def(graph_def, name="")
        g = tf.get_default_graph()
        inp_dic = {}
        for inp in inputs:
            inp_t = g.get_tensor_by_name(inp+":0")
            inp_dic[inp] = inp_t
        out_dic = {}
        for out in outputs:
            out_t = g.get_tensor_by_name(out+":0")
            out_dic[out] = out_t

        sigs[signature_constants.DEFAULT_SERVING_SIGNATURE_DEF_KEY] = \
            tf.saved_model.signature_def_utils.predict_signature_def(
                inp_dic, out_dic)

        builder.add_meta_graph_and_variables(sess,[tag_constants.SERVING],signature_def_map=sigs)
    print("Exporting saved model to:", SAVED_MODEL_DIR+' ...')
    builder.save()

    print("Saved model exported to:", SAVED_MODEL_DIR)
    _show_all(SAVED_MODEL_DIR)
    pb_visual_writer = tf.summary.FileWriter(EXPORT_DIR)
    pb_visual_writer.add_graph(sess.graph)
    print("Visualize the model by running: "
          "tensorboard --logdir={}".format(SAVED_MODEL_DIR))
    with open('/tmp/saved_model_dir.txt', 'w') as f:
        f.write(SAVED_MODEL_DIR)
    with open('/tmp/export_dir.txt', 'w') as f:
        f.write(EXPORT_DIR)
if __name__ == "__main__":
    main()