import React from 'react';
import { connect } from 'react-redux';
import ReactDOM from 'react-dom';
import screenfull from 'screenfull';
import classnames from 'classnames';

import FrameHeader from './FrameHeader';
import FrameQueryEditor from '../containers/FrameQueryEditorContainer';
import {
  FRAME_TYPE_SESSION, FRAME_TYPE_ERROR, FRAME_TYPE_LOADING, FRAME_TYPE_SUCCESS
} from '../lib/const';
import { getShareId } from '../actions';
import { updateFrame } from '../actions/frames';

class FrameLayout extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isFullscreen: false,
      shareId: '',
      shareHidden: false,
      editingQuery: false,
    };
  }

  componentDidUpdate(prevProps, prevState) {
    // If shareId was fetched, select the share url input
    if (prevState.shareId !== this.state.shareId && this.state.shareId !== '') {
      const shareUrlEl = ReactDOM.findDOMNode(this.shareURLEl);
      shareUrlEl.select();
    }
  }

  handleToggleFullscreen = () => {
    if (!screenfull.enabled) {
      return;
    }

    const { isFullscreen } = this.state;

    if (isFullscreen) {
      screenfull.exit();
      this.setState({ isFullscreen: false });
    } else {
      const frameEl = ReactDOM.findDOMNode(this.refs.frame);
      screenfull.request(frameEl);

      // If fullscreen request was successful, set state
      if (screenfull.isFullscreen) {
        this.setState({ isFullscreen: true });
      }
    }
  }

  handleShare = () => {
    const { frame } = this.props;
    const { shareId } = this.state;

    // if shareId is already set, simply toggle the hidden state
    if (shareId) {
      const shareUrlEl = ReactDOM.findDOMNode(this.shareURLEl);

      this.setState({ shareHidden: !this.state.shareHidden });
      shareUrlEl.select();
      return;
    }

    if (frame.type !== FRAME_TYPE_SESSION) {
      return;
    }

    const { query } = frame.data;
    getShareId(query)
      .then(shareId => {
        this.setState({ shareId });
      })
      .catch(err => {
        console.log('error while getting share id', err);
      })
  }

  // saveShareURLRef saves the reference to the share url input as an instance
  // property of this component
  saveShareURLRef = (el) => {
    this.shareURLEl = el;
  }

  // saveCodeMirrorInstance saves the codemirror instance initialized in the
  // <Editor /> component so that we can access it in this component. (e.g. to
  // focus)
  saveCodeMirrorInstance = (cm) => {
    this.queryEditor = cm;
  }

  handleToggleEditingQuery = () => {
    this.setState({
      editingQuery: !this.state.editingQuery
    }, () => {
      if (this.state.editingQuery) {
        this.queryEditor.focus();
      }
    });
  }

  handleToggleCollapse = (done = () => {}) => {
    const { changeCollapseState, frame } = this.props;
    const nextState = !frame.meta.collapsed;

    changeCollapseState(frame, nextState);
    done();
  }

  render() {
    const { children, onDiscardFrame, onSelectQuery, frame } = this.props;
    const { isFullscreen, shareId, shareHidden, editingQuery } = this.state;
    const isCollapsed = frame.meta && frame.meta.collapsed;

    return (
      <li
        className={
          classnames('frame-item', {
            fullscreen: isFullscreen,
            collapsed: isCollapsed,
            'frame-error': frame.type === FRAME_TYPE_ERROR,
            'frame-session': frame.type === FRAME_TYPE_SESSION,
            'frame-loading': frame.type === FRAME_TYPE_LOADING,
            'frame-system': frame.type === FRAME_TYPE_SUCCESS
          })
        }
        ref="frame"
      >
        <FrameHeader
          shareId={shareId}
          onToggleFullscreen={this.handleToggleFullscreen}
          onToggleCollapse={this.handleToggleCollapse}
          onToggleEditingQuery={() => {
            if (frame.meta.collapsed) {
              this.handleToggleCollapse(this.handleToggleEditingQuery);
            } else {
              this.handleToggleEditingQuery();
            }
          }}
          onDiscardFrame={onDiscardFrame}
          onSelectQuery={onSelectQuery}
          onShare={this.handleShare}
          shareHidden={shareHidden}
          frame={frame}
          isFullscreen={isFullscreen}
          isCollapsed={isCollapsed}
          saveShareURLRef={this.saveShareURLRef}
          editingQuery={editingQuery}
        />

        <div className="body-container">
          {isCollapsed ? null : children}
        </div>
      </li>
    );
  }
}

const mapStateToProps = state => ({
});

const mapDispatchToProps = dispatch => ({
  changeCollapseState(frame, nextCollapseState) {
    return dispatch(updateFrame({
      id: frame.id,
      type: frame.type,
      data: frame.data,
      meta: Object.assign({}, frame.meta, { collapsed: nextCollapseState }),
    }))
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(FrameLayout);