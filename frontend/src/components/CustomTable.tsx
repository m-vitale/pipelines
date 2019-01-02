/*
 * Copyright 2018 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import * as React from 'react';
import * as Autosuggest from 'react-autosuggest';
// TODO match and parse required editing the node modules to import correctly.
// Perhaps we should consider using allowSyntheticDefaultImports
// tslint:disable-next-line:no-var-requires
const match = require('autosuggest-highlight/match');
// tslint:disable-next-line:no-var-requires
const parse = require('autosuggest-highlight/parse');
import ArrowRight from '@material-ui/icons/ArrowRight';
import Checkbox, { CheckboxProps } from '@material-ui/core/Checkbox';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';
import CircularProgress from '@material-ui/core/CircularProgress';
import ChipInput, { ChipRendererArgs } from 'material-ui-chip-input';
import IconButton from '@material-ui/core/IconButton';
import MenuItem from '@material-ui/core/MenuItem';
import Radio from '@material-ui/core/Radio';
import Separator from '../atoms/Separator';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import TextField, { TextFieldProps } from '@material-ui/core/TextField';
import Tooltip from '@material-ui/core/Tooltip';
import WarningIcon from '@material-ui/icons/WarningRounded';
import { ListRequest } from '../lib/Apis';
import { classes, stylesheet } from 'typestyle';
import { fonts, fontsize, dimension, commonCss, color, padding } from '../Css';
import { logger } from '../lib/Utils';
import Chip from '@material-ui/core/Chip';
import Paper from '@material-ui/core/Paper';
import { SuggestionsFetchRequestedParams, SuggestionHighlightedParams, SuggestionSelectedEventData } from 'react-autosuggest';
import { InputProps } from '@material-ui/core/Input';

export enum ExpandState {
  COLLAPSED,
  EXPANDED,
  NONE,
}

export interface Column {
  flex?: number;
  label: string;
  sortKey?: string;
  customRenderer?: (value: any, id: string) => React.StatelessComponent;
}

export interface Row {
  expandState?: ExpandState;
  error?: string;
  id: string;
  otherFields: any[];
}

const rowHeight = 40;

export const css = stylesheet({
  cell: {
    $nest: {
      '&:not(:nth-child(2))': {
        color: color.inactive,
      },
    },
    alignSelf: 'center',
    borderBottom: 'initial',
    color: color.foreground,
    fontFamily: fonts.secondary,
    fontSize: fontsize.base,
    letterSpacing: 0.25,
    marginRight: 20,
    overflow: 'hidden',
    textOverflow: 'ellipsis',
    whiteSpace: 'nowrap',
  },
  columnName: {
    color: '#1F1F1F',
    fontSize: fontsize.small,
    fontWeight: 'bold',
    letterSpacing: 0.25,
    marginRight: 20,
  },
  emptyMessage: {
    padding: 20,
    textAlign: 'center',
  },
  expandButton: {
    marginRight: 10,
    padding: 3,
    transition: 'transform 0.3s',
  },
  expandButtonExpanded: {
    transform: 'rotate(90deg)',
  },
  expandableContainer: {
    transition: 'margin 0.2s',
  },
  expandedContainer: {
    borderRadius: 10,
    boxShadow: '0 1px 2px 0 rgba(60,64,67,0.30), 0 1px 3px 1px rgba(60,64,67,0.15)',
    margin: '16px 2px',
  },
  expandedRow: {
    borderBottom: '1px solid transparent !important',
    boxSizing: 'border-box',
    height: '40px !important',
  },
  filterBox: {
    margin: '16px 0'
  },
  filterLabel: {
    transform: 'translate(14px, 26px) scale(1)',
  },
  footer: {
    borderBottom: '1px solid ' + color.divider,
    fontFamily: fonts.secondary,
    height: 40,
    textAlign: 'right',
  },
  header: {
    borderBottom: 'solid 1px ' + color.divider,
    color: color.strong,
    display: 'flex',
    flex: '0 0 40px',
    lineHeight: '40px', // must declare px
  },
  icon: {
    color: color.alert,
    height: 18,
    paddingRight: 4,
    verticalAlign: 'sub',
    width: 18,
  },
  row: {
    $nest: {
      '&:hover': {
        backgroundColor: '#f3f3f3',
      },
    },
    borderBottom: '1px solid #ddd',
    display: 'flex',
    flexShrink: 0,
    height: rowHeight,
    outline: 'none',
  },
  rowsPerPage: {
    color: color.strong,
    height: dimension.xsmall,
    minWidth: dimension.base,
  },
  selected: {
    backgroundColor: color.activeBg,
  },
  selectionToggle: {
    marginRight: 12,
  },
});

interface FilterChip {
  key: string;
  value: {
    type: 'Name' | 'CreatedAt';
    x: number | string;
  };
}

interface CustomTableProps {
  columns: Column[];
  disablePaging?: boolean;
  disableSelection?: boolean;
  disableSorting?: boolean;
  emptyMessage?: string;
  filterString?: string;
  getExpandComponent?: (index: number) => React.ReactNode;
  initialSortColumn?: string;
  initialSortOrder?: 'asc' | 'desc';
  reload: (request: ListRequest) => Promise<string>;
  rows: Row[];
  selectedIds?: string[];
  toggleExpansion?: (rowId: number) => void;
  updateSelection?: (selectedIds: string[]) => void;
  useRadioButtons?: boolean;
}

interface CustomTableState {
  currentPage: number;
  filterBy: string;
  filterChips: FilterChip[];
  filterShouldShowTypes: boolean;
  isBusy: boolean;
  maxPageIndex: number;
  pageSize: number;
  sortBy: string;
  sortOrder: 'asc' | 'desc';
  suggestions: string[];
  textFieldInput: string;
  tokenList: string[];
}

export default class CustomTable extends React.Component<CustomTableProps, CustomTableState> {
  private _isMounted = true;
  private inputRef = React.createRef<React.ComponentType<InputProps>>();

  private FILTER_TYPE_SUGGESTIONS = ['Name', 'Created at'];
  private FILTER_VALUE_SUGGESTIONS = ['Pipeline 1', 'Pipeline 2'];

  constructor(props: CustomTableProps) {
    super(props);

    this.state = {
      currentPage: 0,
      filterBy: '',
      filterChips: [],
      // This will control whether we are showing the filter types (name, timestamp, etc.) or actual values
      filterShouldShowTypes: true,
      isBusy: false,
      maxPageIndex: Number.MAX_SAFE_INTEGER,
      pageSize: 10,
      sortBy: props.initialSortColumn ||
        (props.columns.length ? props.columns[0].sortKey || '' : ''),
      sortOrder: props.initialSortOrder || 'desc',
      suggestions: this.FILTER_TYPE_SUGGESTIONS,
      textFieldInput: 'textFieldInput default value',
      tokenList: [''],
    };
  }

  public handleSelectAllClick(event: React.MouseEvent): void {
    if (this.props.disableSelection === true) {
      return;
    }
    const selectedIds =
      (event.target as CheckboxProps).checked ? this.props.rows.map((v) => v.id) : [];
    if (this.props.updateSelection) {
      this.props.updateSelection(selectedIds);
    }
  }

  public handleClick(e: React.MouseEvent, id: string): void {
    if (this.props.disableSelection === true) {
      return;
    }

    let newSelected = [];
    if (this.props.useRadioButtons) {
      newSelected = [id];
    } else {
      const selectedIds = this.props.selectedIds || [];
      const selectedIndex = selectedIds.indexOf(id);
      newSelected = selectedIndex === -1 ?
        selectedIds.concat(id) :
        selectedIds.slice(0, selectedIndex).concat(selectedIds.slice(selectedIndex + 1));
    }

    if (this.props.updateSelection) {
      this.props.updateSelection(newSelected);
    }

    e.stopPropagation();
  }

  public isSelected(id: string): boolean {
    return !!this.props.selectedIds && this.props.selectedIds.indexOf(id) !== -1;
  }

  public componentDidMount(): void {
    this._pageChanged(0);
  }

  public componentWillUnmount(): void {
    this._isMounted = false;
  }

  public renderChipInput(inputProps: any): JSX.Element {
    // TODO we need to pull 'classes' out of the props, but it conflicts with the 'classes' from type style
    // tslint:disable-next-line:no-shadowed-variable
    const { classes, value, cancelBubble, onChange, onAdd, onDelete, ref, ...other } = inputProps;

    // tslint:disable-next-line:no-console
    console.log('renderChipInput', inputProps);
    // tslint:disable-next-line:no-console
    console.log('value:', value);
    return (
    <ChipInput
      classes={{}}
      InputProps={{
        // this ref is the typical react ref. it is not necessary.
        ref: this.inputRef,
      }}
      // dataSource={[{ key: 'resource_name', value: 'name' }, { key: 'created_at', value: 'CreatedAt' }]}
      // dataSourceConfig={{ text: 'key', value: 'value.x' }}
      value={this.state.filterChips}
      onUpdateInput={onChange}
      onAdd={onAdd}
      onDelete={onDelete}
      // this ref appears to be an empty function, but seems to be necessary
      inputRef={ref}
      {...other}
      chipRenderer={(
        args: ChipRendererArgs,
        key: any,
      ) => {
        // tslint:disable-next-line:no-console
        console.log('ChipRendererArgs', args);
        return (
        <Chip
          key={key}
          style={{
            backgroundColor: args.isFocused ? 'red' : 'green',
            pointerEvents: args.isDisabled ? 'none' : undefined,
          }}
          onClick={args.handleClick}
          onDelete={onDelete}
          label={(args.value as any).value.x}
        />
      );}}
    />);
  }

  public renderSuggestion (suggestion: string, obj: { query: string, isHighlighted: boolean }): JSX.Element {
    // tslint:disable-next-line:no-console
    console.log('renderSuggestion');
    // // tslint:disable-next-line:no-console
    // console.log('suggestion', suggestion);
    // // tslint:disable-next-line:no-console
    // console.log('obj', obj);
    // // tslint:disable-next-line:no-console
    // console.log('-------------------------------');
    const matches = match(suggestion, obj.query);
    const parts = parse(suggestion, matches);

    return (
      <MenuItem
        selected={obj.isHighlighted}
        component='div'
        onMouseDown={(e) => e.preventDefault()} // prevent the click causing the input to be blurred
      >
        <div>
          {parts.map((part: any, index: any) => {
            return part.highlight ? (
              <span key={String(index)} style={{ fontWeight: 500 }}>
                {part.text}
              </span>
            ) : (
              <strong key={String(index)} style={{ fontWeight: 300 }}>
                {part.text}
              </strong>
            );
          })}
        </div>
      </MenuItem>
    );
  }

  public async handleSuggestionsFetchRequested(param: SuggestionsFetchRequestedParams): Promise<void> {
    // tslint:disable-next-line:no-console
    console.log('handleSuggestionsFetchRequested', param);
    // TODO: actually fetch suggestions
    const possibleSuggestions =
      this.state.filterShouldShowTypes
      ? this.FILTER_TYPE_SUGGESTIONS
      : this.FILTER_VALUE_SUGGESTIONS;
    this.setStateSafe({ suggestions: possibleSuggestions.filter(s => param.value && s.startsWith(param.value))});
  }

  public async handleSuggestionsClearRequested(): Promise<void> {
    // TODO: anything else need to happen?
    // tslint:disable-next-line:no-console
    console.log('clear! (actually reseting, not clearing)');
    const possibleSuggestions =
      this.state.filterShouldShowTypes
        ? this.FILTER_TYPE_SUGGESTIONS
        : this.FILTER_VALUE_SUGGESTIONS;
    this.setStateSafe({ suggestions: possibleSuggestions });
  }

  public renderSuggestionsContainer (options: any): JSX.Element {
    // tslint:disable-next-line:no-console
    // console.log('renderSuggestionsContainer', options);

    const { containerProps, children } = options;
  
    return (
      <Paper {...containerProps} square={true} style={{ width: 120 }} >
        {children}
      </Paper>
    );
  }

  public handleSuggestionHighlighted(params: SuggestionHighlightedParams): void {
    // TODO: this will need to be updated if suggestion becomes more than a string.
    // tslint:disable-next-line:no-console
    console.log('handleSuggestionHighlighted', params);
    if (params && params.suggestion) {
      this.setStateSafe({ textFieldInput: params.suggestion });
    }
  }

  public handleTextFieldInputChange = (event: React.FormEvent<any>, param: Autosuggest.ChangeEvent) => {
    // tslint:disable-next-line:no-console
    console.log('handletextFieldInputChange', event, param);
    this.setStateSafe({
      textFieldInput: param.newValue
    });
  };

  public shouldRenderSuggestion(value: string): boolean {
    const { suggestions } = this.state;
    // tslint:disable-next-line:no-console
    // console.log('shouldRenderSuggestions', value, !!value && !!suggestions.find(s => s.toLocaleLowerCase().startsWith(value.toLocaleLowerCase())));
    return !!value && !!suggestions.find(s => s.toLocaleLowerCase().startsWith(value.toLocaleLowerCase()));
  }

  public onSuggestionSelected(event: React.FormEvent<any>, data: SuggestionSelectedEventData<string>): void {
    // tslint:disable-next-line:no-console
    console.log('onSuggestionSelected', this.inputRef.current);
    this.setStateSafe({
      filterShouldShowTypes: !this.state.filterShouldShowTypes,
      textFieldInput: data.suggestion + ':'
    });
    (this.inputRef.current as any).value = data.suggestion + ':';
    this._handleAddChip(data.suggestion + ':');
    return event.preventDefault();
  }

  public render(): JSX.Element {
    const { suggestions, pageSize, sortBy, sortOrder } = this.state;
    // tslint:disable-next-line:no-console
    console.log(suggestions);
    const numSelected = (this.props.selectedIds || []).length;
    const totalFlex = this.props.columns.reduce((total, c) => total += (c.flex || 1), 0);
    const widths = this.props.columns.map(c => (c.flex || 1) / totalFlex * 100);

    return (
      <div className={commonCss.pageOverflowHidden}>

        {/* Filter/Search bar */}
        <div>
          <Autosuggest
            // alwaysRenderSuggestions={true}
            // theme={{
            //   container: classes.container,
            //   suggestionsContainerOpen: classes.suggestionsContainerOpen,
            //   suggestionsList: classes.suggestionsList,
            //   suggestion: classes.suggestion
            // }}
            renderInputComponent={this.renderChipInput.bind(this)}
            shouldRenderSuggestions={this.shouldRenderSuggestion.bind(this)}
            suggestions={suggestions}
            onSuggestionsClearRequested={this.handleSuggestionsClearRequested.bind(this)}
            onSuggestionsFetchRequested={this.handleSuggestionsFetchRequested.bind(this)}
            onSuggestionHighlighted={this.handleSuggestionHighlighted.bind(this)}
            onSuggestionSelected={this.onSuggestionSelected.bind(this)}
            renderSuggestionsContainer={this.renderSuggestionsContainer.bind(this)}
            getSuggestionValue={(suggestion) =>
              // tslint:disable-next-line:no-console
              { console.log('getSuggestionValue'); return suggestion; }}
            renderSuggestion={this.renderSuggestion.bind(this)}
            // focusInputOnSuggestionClick={false}
            inputProps={{
              cancelBubble: true,
              chips: this.state.filterChips,
              classes,
              onAdd: this._handleAddChip.bind(this),
              onChange: this.handleTextFieldInputChange.bind(this),
              onDelete: this._handleDeleteChip.bind(this),
              value: this.state.textFieldInput,
            }}
          />
        </div>
        <br />
        <br />
        <div>Current value of "textFieldInput":</div>
        <div style={{fontFamily: 'monospace'}}>{this.state.textFieldInput}</div>
        <br />
        <div>filterShouldShowTypes:</div>
        <div style={{fontFamily: 'monospace'}}>{'' + this.state.filterShouldShowTypes}</div>
        

        {/* Header */}
        <div className={classes(
          css.header, (this.props.disableSelection || this.props.useRadioButtons) && padding(20, 'l'))}>
          {(this.props.disableSelection !== true && this.props.useRadioButtons !== true) && (
            <div className={classes(css.columnName, css.cell, css.selectionToggle)}>
              <Checkbox indeterminate={!!numSelected && numSelected < this.props.rows.length}
                color='primary' checked={!!numSelected && numSelected === this.props.rows.length}
                onChange={this.handleSelectAllClick.bind(this)} />
            </div>
          )}
          {/* Shift cells to account for expand button */}
          {!!this.props.getExpandComponent && (
            <Separator orientation='horizontal' units={40} />
          )}
          {this.props.columns.map((col, i) => {
            const isColumnSortable = !!this.props.columns[i].sortKey;
            const isCurrentSortColumn = sortBy === this.props.columns[i].sortKey;
            return (
              <div key={i} style={{ width: widths[i] + '%' }}
                className={css.columnName}>
                {this.props.disableSorting === true && <div>{col.label}</div>}
                {!this.props.disableSorting && (
                  <Tooltip title={isColumnSortable ? 'Sort' : 'Cannot sort by this column'}
                    enterDelay={300}>
                    <TableSortLabel active={isCurrentSortColumn} className={commonCss.ellipsis}
                      direction={isColumnSortable ? sortOrder : undefined}
                      onClick={() => this._requestSort(this.props.columns[i].sortKey)}>
                      {col.label}
                    </TableSortLabel>
                  </Tooltip>
                )}
              </div>
            );
          })}
        </div>

        {/* Body */}
        <div className={commonCss.scrollContainer} style={{ minHeight: 60 }}>
          {/* Busy experience */}
          {this.state.isBusy && (<React.Fragment>
            <div className={commonCss.busyOverlay} />
            <CircularProgress size={25} className={commonCss.absoluteCenter} style={{ zIndex: 2 }} />
          </React.Fragment>)}

          {/* Empty experience */}
          {this.props.rows.length === 0 && !!this.props.emptyMessage && !this.state.isBusy && (
            <div className={css.emptyMessage}>{this.props.emptyMessage}</div>
          )}
          {this.props.rows.map((row, i) => {
            if (row.otherFields.length !== this.props.columns.length) {
              logger.error('Rows must have the same number of cells defined in columns');
              return null;
            }
            return (<div className={classes(css.expandableContainer,
              row.expandState === ExpandState.EXPANDED && css.expandedContainer)} key={i}>
              <div role='checkbox' tabIndex={-1} className={
                classes(
                  'tableRow',
                  css.row,
                  this.props.disableSelection === true && padding(20, 'l'),
                  this.isSelected(row.id) && css.selected,
                  row.expandState === ExpandState.EXPANDED && css.expandedRow
                )}
                onClick={e => this.handleClick(e, row.id)}>
                {(this.props.disableSelection !== true || !!this.props.getExpandComponent) && (
                  <div className={classes(css.cell, css.selectionToggle)}>
                    {/* If using checkboxes */}
                    {(this.props.disableSelection !== true && this.props.useRadioButtons !== true) && (
                      <Checkbox color='primary' checked={this.isSelected(row.id)} />)}
                    {/* If using radio buttons */}
                    {(this.props.disableSelection !== true && this.props.useRadioButtons) && (
                      <Radio color='primary' checked={this.isSelected(row.id)} />)}
                    {!!this.props.getExpandComponent && (
                      <IconButton className={classes(css.expandButton,
                        row.expandState === ExpandState.EXPANDED && css.expandButtonExpanded)}
                        onClick={(e) => this._expandButtonToggled(e, i)}>
                        <ArrowRight />
                      </IconButton>
                    )}
                  </div>
                )}
                {row.otherFields.map((cell, c) => (
                  <div key={c} style={{ width: widths[c] + '%' }} className={css.cell}>
                    {c === 0 && row.error && (
                      <Tooltip title={row.error}><WarningIcon className={css.icon} /></Tooltip>
                    )}
                    {this.props.columns[c].customRenderer ?
                      this.props.columns[c].customRenderer!(cell, row.id) : cell}
                  </div>
                ))}
              </div>
              {row.expandState === ExpandState.EXPANDED && this.props.getExpandComponent && (
                <div className={padding(20, 'lrb')}>
                  {this.props.getExpandComponent(i)}
                </div>
              )}
            </div>);
          })}
        </div>

        {/* Footer */}
        {!this.props.disablePaging && (
          <div className={css.footer}>
            <span className={padding(10, 'r')}>Rows per page:</span>
            <TextField select={true} variant='standard' className={css.rowsPerPage}
              InputProps={{ disableUnderline: true }} onChange={this._requestRowsPerPage.bind(this)}
              value={pageSize}>
              {[10, 20, 50, 100].map((size, i) => (
                <MenuItem key={i} value={size}>{size}</MenuItem>
              ))}
            </TextField>

            <IconButton onClick={() => this._pageChanged(-1)} disabled={!this.state.currentPage}>
              <ChevronLeft />
            </IconButton>
            <IconButton onClick={() => this._pageChanged(1)}
              disabled={this.state.currentPage >= this.state.maxPageIndex}>
              <ChevronRight />
            </IconButton>
          </div>
        )}
      </div>
    );
  }

  public async reload(loadRequest?: ListRequest): Promise<string> {
    // Override the current state with incoming request
    const request: ListRequest = Object.assign({
      filterBy: this.state.filterBy,
      orderAscending: this.state.sortOrder === 'asc',
      pageSize: this.state.pageSize,
      pageToken: this.state.tokenList[this.state.currentPage],
      sortBy: this.state.sortBy,
    }, loadRequest);

    let result = '';
    try {
      this.setStateSafe({
        filterBy: request.filterBy,
        isBusy: true,
        pageSize: request.pageSize!,
        sortBy: request.sortBy!,
        sortOrder: request.orderAscending ? 'asc' : 'desc',
      });

      if (request.sortBy && !request.orderAscending) {
        request.sortBy += ' desc';
      }

      result = await this.props.reload(request);
    } finally {
      this.setStateSafe({ isBusy: false });
    }
    return result;
  }

  // private _handleChange = (name: string) => (event: any) => {
  //   // TODO: add rate-limiting here, or within _requestFilter()
  //   // tslint:disable-next-line:no-console
  //   console.log('change: ', name);
  //   const value = (event.target as TextFieldProps).value;
  //   this.setStateSafe({ [name]: value } as any, this._requestFilter.bind(this)(value));
  // }

  
  private _handleAddChip(chip: string): void {
    // tslint:disable-next-line:no-console
    console.log('_handleAddChip', chip);
    // tslint:disable-next-line:no-console
    console.log(JSON.stringify(chip));
    // tslint:disable-next-line:no-console
    console.log(chip);
    const chips = this.state.filterChips;
    // TODO: lots of work needed here.
    chips.push({ key: chips.length + '', value: { type: 'Name', x: `${chip}` } });
    // tslint:disable-next-line:no-console
    console.log(chips);
    // this.setStateSafe({ filterChips: chips, textFieldInput: '' }, async () => this._requestFilter());
    this.setStateSafe({ filterChips: chips }, async () => this._requestFilter());
  }

  private _handleDeleteChip(chip: FilterChip, index: number): void {
    // tslint:disable-next-line:no-console
    console.log('_handleDeleteChip', chip, index);
    const chips = this.state.filterChips;
    chips.splice(index, 1);
    this.setStateSafe({ filterChips: chips }, async () => this._requestFilter());
  }
  

  private async _requestFilter(): Promise<void> {
    if (this.state.filterChips.length) {
      this._resetToFirstPage(
        // TODO: convert this to actual filter, not just JSON string.
        await this.reload({ filterBy: JSON.stringify(this.state.filterChips) })
      );
    }
  }

  private setStateSafe(newState: Partial<CustomTableState>, cb?: () => void): void {
    if (this._isMounted) {
      this.setState(newState as any, cb);
    }
  }

  private _requestSort(sortBy?: string): void {
    if (sortBy) {
      // Set the sort column to the provided column if it's different, and
      // invert the sort order it if it's the same column
      const sortOrder = this.state.sortBy === sortBy ?
        (this.state.sortOrder === 'asc' ? 'desc' : 'asc') : 'asc';
      this.setStateSafe({ sortOrder, sortBy }, async () => {
        this._resetToFirstPage(
          await this.reload({ pageToken: '', orderAscending: sortOrder === 'asc', sortBy }));
      });
    }
  }

  private async _pageChanged(offset: number): Promise<void> {
    let newCurrentPage = this.state.currentPage + offset;
    let maxPageIndex = this.state.maxPageIndex;
    newCurrentPage = Math.max(0, newCurrentPage);
    newCurrentPage = Math.min(this.state.maxPageIndex, newCurrentPage);

    const newPageToken = await this.reload({
      pageToken: this.state.tokenList[newCurrentPage],
    });

    if (newPageToken) {
      // If we're using the greatest yet known page, then the pageToken will be new.
      if (newCurrentPage + 1 === this.state.tokenList.length) {
        this.state.tokenList.push(newPageToken);
      }
    } else {
      maxPageIndex = newCurrentPage;
    }

    this.setStateSafe({ currentPage: newCurrentPage, maxPageIndex });
  }

  private async _requestRowsPerPage(event: React.ChangeEvent): Promise<void> {
    const pageSize = (event.target as TextFieldProps).value as number;

    this._resetToFirstPage(await this.reload({ pageSize, pageToken: '' }));
  }

  private _resetToFirstPage(newPageToken?: string): void {
    let maxPageIndex = Number.MAX_SAFE_INTEGER;
    const newTokenList = [''];

    if (newPageToken) {
      newTokenList.push(newPageToken);
    } else {
      maxPageIndex = 0;
    }

    // Reset state, since this invalidates the token list and page counter calculations
    this.setStateSafe({
      currentPage: 0,
      maxPageIndex,
      tokenList: newTokenList,
    });
  }

  private _expandButtonToggled(e: React.MouseEvent, rowIndex: number): void {
    e.stopPropagation();
    if (this.props.toggleExpansion) {
      this.props.toggleExpansion(rowIndex);
    }
  }
}
