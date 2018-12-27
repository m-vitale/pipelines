/// <reference path="./custom.d.ts" />
// tslint:disable
/**
 * backend/api/experiment.proto
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: version not set
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


import * as url from "url";
import * as portableFetch from "portable-fetch";
import { Configuration } from "./configuration";

const BASE_PATH = "http://localhost".replace(/\/+$/, "");

/**
 *
 * @export
 */
export const COLLECTION_FORMATS = {
    csv: ",",
    ssv: " ",
    tsv: "\t",
    pipes: "|",
};

/**
 *
 * @export
 * @interface FetchAPI
 */
export interface FetchAPI {
    (url: string, init?: any): Promise<Response>;
}

/**
 *  
 * @export
 * @interface FetchArgs
 */
export interface FetchArgs {
    url: string;
    options: any;
}

/**
 * 
 * @export
 * @class BaseAPI
 */
export class BaseAPI {
    protected configuration: Configuration;

    constructor(configuration?: Configuration, protected basePath: string = BASE_PATH, protected fetch: FetchAPI = portableFetch) {
        if (configuration) {
            this.configuration = configuration;
            this.basePath = configuration.basePath || this.basePath;
        }
    }
};

/**
 * 
 * @export
 * @class RequiredError
 * @extends {Error}
 */
export class RequiredError extends Error {
    name: "RequiredError"
    constructor(public field: string, msg?: string) {
        super(msg);
    }
}

/**
 * 
 * @export
 * @interface ApiExperiment
 */
export interface ApiExperiment {
    /**
     * Output. Unique experiment ID. Generated by API server.
     * @type {string}
     * @memberof ApiExperiment
     */
    id?: string;
    /**
     * Required input field. Unique experiment name provided by user.
     * @type {string}
     * @memberof ApiExperiment
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ApiExperiment
     */
    description?: string;
    /**
     * Output. The time that the experiment created.
     * @type {Date}
     * @memberof ApiExperiment
     */
    created_at?: Date;
}

/**
 * Filter is used to filter resources returned from a ListXXX request.  Example filters: 1) Filter runs with status = 'Running' filter {   predicate {     key: \"status\"     op: EQUALS     string_value: \"Running\"   } }  2) Filter runs that succeeded since Dec 1, 2018 filter {   predicate {     key: \"status\"     op: EQUALS     string_value: \"Succeeded\"   }   predicate {     key: \"created_at\"     op: GREATER_THAN     timestamp_value {       seconds: 1543651200     }   } }  3) Filter runs with one of labels 'label_1' or 'label_2'  filter {   predicate {     key: \"label\"     op: IN     string_values {       value: 'label_1'       value: 'label_2'     }   } }
 * @export
 * @interface ApiFilter
 */
export interface ApiFilter {
    /**
     * All predicates are AND-ed when this filter is applied.
     * @type {Array&lt;ApiPredicate&gt;}
     * @memberof ApiFilter
     */
    predicates?: Array<ApiPredicate>;
}

/**
 * 
 * @export
 * @interface ApiIntValues
 */
export interface ApiIntValues {
    /**
     * 
     * @type {Array&lt;number&gt;}
     * @memberof ApiIntValues
     */
    values?: Array<number>;
}

/**
 * 
 * @export
 * @interface ApiListExperimentsResponse
 */
export interface ApiListExperimentsResponse {
    /**
     * 
     * @type {Array&lt;ApiExperiment&gt;}
     * @memberof ApiListExperimentsResponse
     */
    experiments?: Array<ApiExperiment>;
    /**
     * 
     * @type {number}
     * @memberof ApiListExperimentsResponse
     */
    total_size?: number;
    /**
     * 
     * @type {string}
     * @memberof ApiListExperimentsResponse
     */
    next_page_token?: string;
}

/**
 * 
 * @export
 * @interface ApiLongValues
 */
export interface ApiLongValues {
    /**
     * 
     * @type {Array&lt;string&gt;}
     * @memberof ApiLongValues
     */
    values?: Array<string>;
}

/**
 * Predicate captures individual conditions that must be true for a resource being filtered.
 * @export
 * @interface ApiPredicate
 */
export interface ApiPredicate {
    /**
     * 
     * @type {PredicateOp}
     * @memberof ApiPredicate
     */
    op?: PredicateOp;
    /**
     * 
     * @type {string}
     * @memberof ApiPredicate
     */
    key?: string;
    /**
     * 
     * @type {number}
     * @memberof ApiPredicate
     */
    int_value?: number;
    /**
     * 
     * @type {string}
     * @memberof ApiPredicate
     */
    long_value?: string;
    /**
     * 
     * @type {string}
     * @memberof ApiPredicate
     */
    string_value?: string;
    /**
     * Timestamp values will be converted to Unix time (seconds since the epoch) prior to being used in a filtering operation.
     * @type {Date}
     * @memberof ApiPredicate
     */
    timestamp_value?: Date;
    /**
     * Array values below are only meant to be used by the IN operator.
     * @type {ApiIntValues}
     * @memberof ApiPredicate
     */
    int_values?: ApiIntValues;
    /**
     * 
     * @type {ApiLongValues}
     * @memberof ApiPredicate
     */
    long_values?: ApiLongValues;
    /**
     * 
     * @type {ApiStringValues}
     * @memberof ApiPredicate
     */
    string_values?: ApiStringValues;
}

/**
 * 
 * @export
 * @interface ApiStatus
 */
export interface ApiStatus {
    /**
     * 
     * @type {string}
     * @memberof ApiStatus
     */
    error?: string;
    /**
     * 
     * @type {number}
     * @memberof ApiStatus
     */
    code?: number;
    /**
     * 
     * @type {Array&lt;ProtobufAny&gt;}
     * @memberof ApiStatus
     */
    details?: Array<ProtobufAny>;
}

/**
 * 
 * @export
 * @interface ApiStringValues
 */
export interface ApiStringValues {
    /**
     * 
     * @type {Array&lt;string&gt;}
     * @memberof ApiStringValues
     */
    values?: Array<string>;
}

/**
 * Op is the operation to apply.   - EQUALS: Operators on scalar values. Only applies to one of |int_value|, |long_value|, |string_value| or |timestamp_value|.  - IN: Checks if the value is a member of a given array, which should be one of |int_values|, |long_values| or |string_values|.
 * @export
 * @enum {string}
 */
export enum PredicateOp {
    UNKNOWN = <any> 'UNKNOWN',
    EQUALS = <any> 'EQUALS',
    NOTEQUALS = <any> 'NOT_EQUALS',
    GREATERTHAN = <any> 'GREATER_THAN',
    GREATERTHANEQUALS = <any> 'GREATER_THAN_EQUALS',
    LESSTHAN = <any> 'LESS_THAN',
    LESSTHANEQUALS = <any> 'LESS_THAN_EQUALS',
    IN = <any> 'IN'
}

/**
 * `Any` contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message.  Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type.  Example 1: Pack and unpack a message in C++.      Foo foo = ...;     Any any;     any.PackFrom(foo);     ...     if (any.UnpackTo(&foo)) {       ...     }  Example 2: Pack and unpack a message in Java.      Foo foo = ...;     Any any = Any.pack(foo);     ...     if (any.is(Foo.class)) {       foo = any.unpack(Foo.class);     }   Example 3: Pack and unpack a message in Python.      foo = Foo(...)     any = Any()     any.Pack(foo)     ...     if any.Is(Foo.DESCRIPTOR):       any.Unpack(foo)       ...   Example 4: Pack and unpack a message in Go       foo := &pb.Foo{...}      any, err := ptypes.MarshalAny(foo)      ...      foo := &pb.Foo{}      if err := ptypes.UnmarshalAny(any, foo); err != nil {        ...      }  The pack methods provided by protobuf library will by default use 'type.googleapis.com/full.type.name' as the type URL and the unpack methods only use the fully qualified type name after the last '/' in the type URL, for example \"foo.bar.com/x/y.z\" will yield type name \"y.z\".   JSON ==== The JSON representation of an `Any` value uses the regular representation of the deserialized, embedded message, with an additional field `@type` which contains the type URL. Example:      package google.profile;     message Person {       string first_name = 1;       string last_name = 2;     }      {       \"@type\": \"type.googleapis.com/google.profile.Person\",       \"firstName\": <string>,       \"lastName\": <string>     }  If the embedded message type is well-known and has a custom JSON representation, that representation will be embedded adding a field `value` which holds the custom JSON in addition to the `@type` field. Example (for message [google.protobuf.Duration][]):      {       \"@type\": \"type.googleapis.com/google.protobuf.Duration\",       \"value\": \"1.212s\"     }
 * @export
 * @interface ProtobufAny
 */
export interface ProtobufAny {
    /**
     * A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one \"/\" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading \".\" is not accepted).  In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows:  * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a [google.protobuf.Type][]   value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the   URL, or have them precompiled into a binary to avoid any   lookup. Therefore, binary compatibility needs to be preserved   on changes to types. (Use versioned type names to manage   breaking changes.)  Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com.  Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics.
     * @type {string}
     * @memberof ProtobufAny
     */
    type_url?: string;
    /**
     * Must be a valid serialized protocol buffer of the above specified type.
     * @type {string}
     * @memberof ProtobufAny
     */
    value?: string;
}


/**
 * ExperimentServiceApi - fetch parameter creator
 * @export
 */
export const ExperimentServiceApiFetchParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @param {ApiExperiment} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createExperiment(body: ApiExperiment, options: any = {}): FetchArgs {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling createExperiment.');
            }
            const localVarPath = `/apis/v1beta1/experiments`;
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'POST' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication Bearer required
            if (configuration && configuration.apiKey) {
                const localVarApiKeyValue = typeof configuration.apiKey === 'function'
					? configuration.apiKey("authorization")
					: configuration.apiKey;
                localVarHeaderParameter["authorization"] = localVarApiKeyValue;
            }

            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);
            const needsSerialization = (<any>"ApiExperiment" !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.body =  needsSerialization ? JSON.stringify(body || {}) : (body || "");

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteExperiment(id: string, options: any = {}): FetchArgs {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteExperiment.');
            }
            const localVarPath = `/apis/v1beta1/experiments/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'DELETE' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication Bearer required
            if (configuration && configuration.apiKey) {
                const localVarApiKeyValue = typeof configuration.apiKey === 'function'
					? configuration.apiKey("authorization")
					: configuration.apiKey;
                localVarHeaderParameter["authorization"] = localVarApiKeyValue;
            }

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getExperiment(id: string, options: any = {}): FetchArgs {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getExperiment.');
            }
            const localVarPath = `/apis/v1beta1/experiments/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'GET' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication Bearer required
            if (configuration && configuration.apiKey) {
                const localVarApiKeyValue = typeof configuration.apiKey === 'function'
					? configuration.apiKey("authorization")
					: configuration.apiKey;
                localVarHeaderParameter["authorization"] = localVarApiKeyValue;
            }

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} [page_token] 
         * @param {number} [page_size] 
         * @param {string} [sort_by] Can be format of \&quot;field_name\&quot;, \&quot;field_name asc\&quot; or \&quot;field_name des\&quot; Ascending by default.
         * @param {string} [filter] A base-64 encoded, JSON-serialized Filter protocol buffer (see filter.proto).
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listExperiment(page_token?: string, page_size?: number, sort_by?: string, filter?: string, options: any = {}): FetchArgs {
            const localVarPath = `/apis/v1beta1/experiments`;
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'GET' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication Bearer required
            if (configuration && configuration.apiKey) {
                const localVarApiKeyValue = typeof configuration.apiKey === 'function'
					? configuration.apiKey("authorization")
					: configuration.apiKey;
                localVarHeaderParameter["authorization"] = localVarApiKeyValue;
            }

            if (page_token !== undefined) {
                localVarQueryParameter['page_token'] = page_token;
            }

            if (page_size !== undefined) {
                localVarQueryParameter['page_size'] = page_size;
            }

            if (sort_by !== undefined) {
                localVarQueryParameter['sort_by'] = sort_by;
            }

            if (filter !== undefined) {
                localVarQueryParameter['filter'] = filter;
            }

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * ExperimentServiceApi - functional programming interface
 * @export
 */
export const ExperimentServiceApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @param {ApiExperiment} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createExperiment(body: ApiExperiment, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<ApiExperiment> {
            const localVarFetchArgs = ExperimentServiceApiFetchParamCreator(configuration).createExperiment(body, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteExperiment(id: string, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<any> {
            const localVarFetchArgs = ExperimentServiceApiFetchParamCreator(configuration).deleteExperiment(id, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getExperiment(id: string, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<ApiExperiment> {
            const localVarFetchArgs = ExperimentServiceApiFetchParamCreator(configuration).getExperiment(id, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
        /**
         * 
         * @param {string} [page_token] 
         * @param {number} [page_size] 
         * @param {string} [sort_by] Can be format of \&quot;field_name\&quot;, \&quot;field_name asc\&quot; or \&quot;field_name des\&quot; Ascending by default.
         * @param {string} [filter] A base-64 encoded, JSON-serialized Filter protocol buffer (see filter.proto).
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listExperiment(page_token?: string, page_size?: number, sort_by?: string, filter?: string, options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<ApiListExperimentsResponse> {
            const localVarFetchArgs = ExperimentServiceApiFetchParamCreator(configuration).listExperiment(page_token, page_size, sort_by, filter, options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
    }
};

/**
 * ExperimentServiceApi - factory interface
 * @export
 */
export const ExperimentServiceApiFactory = function (configuration?: Configuration, fetch?: FetchAPI, basePath?: string) {
    return {
        /**
         * 
         * @param {ApiExperiment} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createExperiment(body: ApiExperiment, options?: any) {
            return ExperimentServiceApiFp(configuration).createExperiment(body, options)(fetch, basePath);
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteExperiment(id: string, options?: any) {
            return ExperimentServiceApiFp(configuration).deleteExperiment(id, options)(fetch, basePath);
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getExperiment(id: string, options?: any) {
            return ExperimentServiceApiFp(configuration).getExperiment(id, options)(fetch, basePath);
        },
        /**
         * 
         * @param {string} [page_token] 
         * @param {number} [page_size] 
         * @param {string} [sort_by] Can be format of \&quot;field_name\&quot;, \&quot;field_name asc\&quot; or \&quot;field_name des\&quot; Ascending by default.
         * @param {string} [filter] A base-64 encoded, JSON-serialized Filter protocol buffer (see filter.proto).
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listExperiment(page_token?: string, page_size?: number, sort_by?: string, filter?: string, options?: any) {
            return ExperimentServiceApiFp(configuration).listExperiment(page_token, page_size, sort_by, filter, options)(fetch, basePath);
        },
    };
};

/**
 * ExperimentServiceApi - object-oriented interface
 * @export
 * @class ExperimentServiceApi
 * @extends {BaseAPI}
 */
export class ExperimentServiceApi extends BaseAPI {
    /**
     * 
     * @param {} body 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExperimentServiceApi
     */
    public createExperiment(body: ApiExperiment, options?: any) {
        return ExperimentServiceApiFp(this.configuration).createExperiment(body, options)(this.fetch, this.basePath);
    }

    /**
     * 
     * @param {} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExperimentServiceApi
     */
    public deleteExperiment(id: string, options?: any) {
        return ExperimentServiceApiFp(this.configuration).deleteExperiment(id, options)(this.fetch, this.basePath);
    }

    /**
     * 
     * @param {} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExperimentServiceApi
     */
    public getExperiment(id: string, options?: any) {
        return ExperimentServiceApiFp(this.configuration).getExperiment(id, options)(this.fetch, this.basePath);
    }

    /**
     * 
     * @param {} [page_token] 
     * @param {} [page_size] 
     * @param {} [sort_by] Can be format of \&quot;field_name\&quot;, \&quot;field_name asc\&quot; or \&quot;field_name des\&quot; Ascending by default.
     * @param {} [filter] A base-64 encoded, JSON-serialized Filter protocol buffer (see filter.proto).
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExperimentServiceApi
     */
    public listExperiment(page_token?: string, page_size?: number, sort_by?: string, filter?: string, options?: any) {
        return ExperimentServiceApiFp(this.configuration).listExperiment(page_token, page_size, sort_by, filter, options)(this.fetch, this.basePath);
    }

}

