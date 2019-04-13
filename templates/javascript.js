var log = function(msg) {
  console.log(msg);
};
var debug = function(msg) {
  if (false) log(msg);
};
// short for document.getElementById(id)
var e = function(id) {
  return document.getElementById(id);
};

// https://stackoverflow.com/questions/7444451/how-to-get-the-actual-rendered-font-when-its-not-defined-in-css
var css = function(element, property) {
  return window.getComputedStyle(element, null).getPropertyValue(property);
};

// https://stackoverflow.com/questions/118241/calculate-text-width-with-javascript/21015393#21015393
var getTextWidth = function(text, fontSize) {
  // if given, use cached canvas for better performance
  // else, create new canvas
  var canvas =
    getTextWidth.canvas ||
    (getTextWidth.canvas = document.createElement("canvas"));
  var fontFamily =
    getTextWidth.fontFamily ||
    (getTextWidth.fontFamily = css(e("body"), "font-family"));
  var context = canvas.getContext("2d");
  context.font = fontSize + "px " + fontFamily;
  var metrics = context.measureText(text);
  return metrics.width;
};
// short for document.getElementById(id).classList.add(class)
var addcls = function(el, cls) {
  el.classList.add(cls);
};
// short for document.getElementById(id).classList.rm(class)
var rmcls = function(el, cls) {
  el.classList.remove(cls);
};
var hide = function(el) {
  addcls(el, "hide");
};
var show = function(el) {
  rmcls(el, "hide");
};
var disable = function(el) {
  el.disabled = "disabled";
};
var enable = function(el) {
  el.disabled = "";
};
var hideId = function(id) {
  hide(e(id));
};
var showId = function(id) {
  show(e(id));
};
var disableId = function(id) {
  disable(e(id));
};
var enableId = function(id) {
  enable(e(id));
};

/*
 * switch betwwen different views
 */
const views = {
  // view name: button ID
  viewPlaylist: "list",
  viewSearch: "search",
  viewDirectory: "browser"
};
var showView = function(view) {
  return function() {
    log("handle:" + view);
    for (var k in views) {
      // disable/enable view buttons & show/hide view
      const el = e(k);
      switch (k) {
        case view: // show matching view
          log("show:" + k);
          show(el);
          disableId(views[k]);
          break;
        default:
          // hidde others
          log("hidde:" + k);
          hide(el);
          enableId(views[k]);
          break;
      }
    }
    // special actions based on the select view
    switch (view) {
      case "viewDirectory":
        command("browse", ""); // send a command
        break;
      case "viewSearch":
        if (view == "viewSearch") {
          const el = e("searchText");
          el.select();
          el.focus();
        }
        break;
    }
  };
};

var addEvent = function(el, type, fn) {
  if (el.addEventListener) {
    el.addEventListener(type, fn, false);
  } else {
    el.attachEvent("on" + type, fn);
  }
};

// Resize voodoo
var resize = function(el, minFS, maxFS) {
  var fs = maxFS;
  var txt = el.innerHTML;
  var maxWidth = el.clientWidth;
  while (fs >= minFS && getTextWidth(txt, fs) + fs > maxWidth) {
    fs -= 2;
  }
  if (fs < minFS) {
    fs = minFS;
    maxWidth *= 2; // we allow one page break
    var truncated = false;
    while (txt.length > 8 && getTextWidth(txt, fs) + fs > maxWidth) {
      log("too long: " + txt);
      txt = txt.slice(0, txt.length - 1);
      truncated = true;
    }
    if (truncated) {
      el.innerHTML = txt.slice(0, txt.length - 3) + "&hellip;";
    }
  }
  el.style.fontSize = fs + "px";
};
var resizer = function () {
  return true;
};
/*function() {
  var el = document.getElementsByClassName("resize");
  var i;
  for (i = 0; i < el.length; i++) {
    resize(el[i], 12, 28);
  }
};*/

var triggerResize = function() {
  window.dispatchEvent(new Event("resize")); // trigger resize events on the song stuff
};

var ws_addr, ws;

// send a command on the websocket
var command = function(cmd, data) {
  if (!ws) return false;
  log({ cmd, data });
  ws.send(JSON.stringify({ command: cmd, data: data }));
  return true;
};

// wrapper to use command as an onclick
// example: e("play").onclick = btnCommand("play", 1)
var btnCommand = function(cmd, data) {
  return function() {
    return command(cmd, data);
  };
};

// a few 'globals' to track the process and current state of the player
var gDuration = 1.0;
var gElapsed = 0.0;
var gState = {
  play: "pause",
  consume: false,
  repeat: false,
  single: false,
  random: false,
  song: 0,
  nextsong: 0
};
var gErrors = [];
var gPlaylistFiles = [];
var readableSeconds = function(value) {
  var min = parseInt(value / 60);
  var sec = parseInt(value % 60);
  if (sec < 10) {
    sec = "0" + sec;
  }
  return min + ":" + sec;
};

// this functions runs forever and gets called every second to update the
// elapsed and duration information of the active song
var updateProgress = function() {
  e("elapsed").innerHTML = readableSeconds(gElapsed);
  e("duration").innerHTML = readableSeconds(gDuration);
  if (gState["play"] == "play" && gElapsed < gDuration) {
    // increment the seconds if playing and not finished
    gElapsed += 1;
  }
  if (gState["play"] == "stop") {
    e("progress").style.width = "0%";
  } else {
    e("progress").style.width = (gElapsed / gDuration) * 100 + "%";
  }
  setTimeout(updateProgress, 1000);
};

var showError = function(msg) {
  const el = e("mainError");
  if (msg != "") {
    gErrors.push(msg);
  }
  // FIXME: looks ugly
  var sep = "";
  var str = "";
  gErrors.map(function(value) {
    str += sep + value;
    sep = "<br />";
  });
  el.innerHTML = str;
  show(el);

  setTimeout(function() {
    if (gErrors.length > 0) {
      gErrors.shift();
      showError("");
    }
    if (gErrors.length == 0) {
      hideError();
    }
  }, 5000);
};

var hideError = function() {
  const el = e("mainError");
  el.innerHTML = "";
  hide(el);
};

var updateStatus = function(data) {
  debug("updateStatus");
  // FIXME: use nextsong (or nextsongid) to highlight the next song
  const {
    state,
    duration,
    elapsed,
    consume,
    repeat,
    random,
    single,
    song,
    nextsong
  } = data;
  gState = {
    play: state,
    consume: consume,
    repeat: repeat,
    single: single,
    random: random,
    song: song,
    nextsong: nextsong
  };
  // debug(`updateStatus(${state})`);
  switch (state) {
    case "pause":
    case "play":
      togglePlayPause(state);
      gDuration = duration;
      gElapsed = elapsed;
      break;
    case "stop":
      stop();
      gDuration = 0.0;
      gElapsed = 0.0;
      break;
  }
  // update the mode ctrl
  ["consume", "repeat", "single", "random"].map(function(value) {
    const eDis = e(value + "Disable");
    const eEn = e(value + "Enable");
    if (gState[value]) {
      show(eDis);
      hide(eEn);
    } else {
      hide(eDis);
      show(eEn);
    }
  });
};

var updateActiveSong = function(data) {
  const { file, artist, title, album_artist, album, position } = data;
  e("ctrlSong").title = file;
  e("artist").innerHTML = artist;
  e("title").innerHTML = title;
  e("album").innerHTML =
    album + (artist != album_artist ? "<wbr>[" + album_artist + "]" : "");
  const el = e("remove");
  if (position == -1) {
    disable(el);
  } else {
    enable(el);
    el.onclick = btnCommand("remove", position.toString());
  }
};

var newSongNode = function(id, entry, nr) {
  const { file, artist, title, album, album_artist, duration } = entry;
  var node = e(id).cloneNode(true);
  node.classList.remove("hide");
  node.title = file;
  node.id = id + nr;
  node.querySelector("#songCellArtist").innerHTML = artist;
  node.querySelector("#songCellTitle").innerHTML = title;
  // only show album artist if it's different
  node.querySelector("#songCellAlbum").innerHTML =
    album + (artist != album_artist ? "<wbr>[" + album_artist + "]" : "");
  node.querySelector("#songCellDuration").innerHTML = readableSeconds(duration);
  return node;
};

var processResponse = function(obj) {
  const { type, data } = obj;
  log({ type, data });

  switch (type) {
    case "error":
    case "info":
      showError(data);
      break;
    case "version":
      log("MPD protocol version: " + data);
      break;
    case "update":
      updateStatus(data.status);

      e("playlist").innerHTML = ""; // delete old playlist
      gPlaylistFiles.length = 0;
      // used to figure out if we need to show prio buttons or not:
      // if state == stop && activeFile is set && queuesize <=2 → no need to show buttons
      // if state == stop && activeFile is not set → show buttons
      // if state != stop && queuesize <=2 → do not show buttons
      // FIXME: is that correct?
      const queueSize = data.queue.length;
      const state = data.status.state;
      const activeFile = data.activeSong.file;
      data.queue.map(function(entry, i) {
        const { file, prio, position, isActive, isNext } = entry;
        if (isActive) {
          return; // don't show the active song in the list
        }
        gPlaylistFiles.push(file);
        var node = newSongNode("playlistEntry", entry, i);
        const btnPlay = node.querySelector("#plPlay");
        if (isActive) {
          disable(btnPlay);
          addcls(node, "activeSong");
        } else {
          btnPlay.onclick = btnCommand("play", position.toString());
        }
        if (isNext) {
          addcls(node, "nextSong");
        }
        node.querySelector("#plRemove").onclick = btnCommand(
          "remove",
          position.toString()
        );
        // FIXME: this only makes sense in certain modes (random + ?)
        const btnPrio1 = node.querySelector("#plPrio1");
        const btnPrio2 = node.querySelector("#plPrio2");
        const btnPrio3 = node.querySelector("#plPrio3");
        if (!gState["random"]) {
          hide(btnPrio1);
          hide(btnPrio2);
          hide(btnPrio3);
        } else {
          // see explanation above
          if (
            (state != "stop" || activeFile != "") &&
            (isActive || queueSize <= 2)
          ) {
            disable(btnPrio1);
            disable(btnPrio2);
            disable(btnPrio3);
          } else {
            switch (prio) {
              case 0:
                disable(btnPrio3);
                break;
              case 127:
                disable(btnPrio2);
                break;
              case 255:
                disable(btnPrio1);
                break;
            }
            btnPrio1.onclick = btnCommand("prio", "255:" + position.toString());
            btnPrio2.onclick = btnCommand("prio", "127:" + position.toString());
            btnPrio3.onclick = btnCommand("prio", "0:" + position.toString());
          }
        }
        e("playlist").append(node);
      });

      updateActiveSong(data.activeSong);
      triggerResize();
      break;
    case "searchResult":
      e("searchResult").innerHTML = ""; // delete old search result
      // debug({ gFiles: gPlaylistFiles });
      if (data.truncated) {
        showError("search result limited to " + data.maxResults);
      }
      if (data.searchResult.length == 0) {
        showError("nothing found");
      }
      data.searchResult.map(function(entry, i) {
        const { file } = entry;
        const node = newSongNode("searchEntry", entry, i);
        const btn = node.querySelector("#srAdd");
        // disable button for files already in the playlist
        if (gPlaylistFiles.includes(file)) {
          // FIXME: should we add a button to remove it from the playlist?
          disable(btn);
        }
        btn.onclick = function() {
          disable(btn);
          return command("add", file);
        };
        e("searchResult").append(node);
      });
      triggerResize();

      break;
    case "directoryList":
      e("directoryList").innerHTML = "";
      if (data.hasParent) {
        node = e("directoryListEntry").cloneNode(true);
        node.id = "dlRowParent";
        node.title = data.parent;
        show(node);
        node.querySelector("#directoryName").innerHTML =
          data.parent != "" ? data.parent : "..";

        ["#directoryName", "#dlBrowse"].map(function(v) {
          node.querySelector(v).onclick = btnCommand("browse", data.parent);
        });
        disable(node.querySelector("#dlAdd"));
        e("directoryList").append(node);
      }
      data.directoryList.map(function(entry, i) {
        const { type, directory, file } = entry;
        var node;
        if (type == "directory") {
          node = e("directoryListEntry").cloneNode(true);
          node.id = "dlRow" + i;
          rmcls(node, "hide");
          node.querySelector("#directoryName").innerHTML = directory;
          ["#directoryName", "#dlBrowse"].map(function(v) {
            node.querySelector(v).onclick = btnCommand("browse", directory);
          });
          const btn = node.querySelector("#dlAdd");
          btn.onclick = function() {
            disable(btn);
            return command("add", directory);
          };
        } else {
          node = newSongNode("searchEntry", entry, i);
          const btn = node.querySelector("#srAdd");
          // disable button for files already in the playlist
          if (gPlaylistFiles.includes(file)) {
            // FIXME: should we add a button to remove it from the playlist?
            disable(btn);
          }
          btn.onclick = function() {
            disable(btn);
            return command("add", file);
          };
        }
        e("directoryList").append(node);
        triggerResize();
      });
      break;
  }
};

var shouldRefresh = false;
var openWebSocket = function() {
  ws = new WebSocket(ws_addr);
  ws.onopen = function(evt) {
    debug("OPEN");
    hideError();
    if (shouldRefresh) {
      location.reload(); // force refresh to get latest html etc
    }
  };
  ws.onclose = function(evt) {
    debug("CLOSE");
    ws = null;
    showError("no connection&hellip; reconnecting in 4 seconds&hellip;");
    shouldRefresh = true;
    setTimeout(openWebSocket, 4000);
  };
  ws.onmessage = function(evt) {
    processResponse(JSON.parse(evt.data));
  };
  ws.onerror = function(evt) {
    log({ evt });
  };
};

// FIXME: play / pause / stop doesn't work as expected (does not enable buttons all the time)
var stop = function() {
  debug("stop");
  ["pause", "resume", "stop", "next", "previous"].map(hideId);
  ["play", "remove"].map(showId);
  ["play", "remove"].map(enableId);
};
var play = function() {
  debug("play");
  ["resume", "play", "remove"].map(hideId);
  ["pause", "stop", "next", "previous"].map(showId);
  ["pause", "stop", "next", "previous"].map(enableId);
};
var pause = function() {
  debug("pause");
  ["pause", "play", "remove"].map(hideId);
  ["resume", "stop", "next", "previous"].map(showId);
  ["resume", "stop", "next", "previous"].map(enableId);
};
var togglePlayPause = function(state) {
  // debug(`togglePlayPause(${state})`);
  switch (state) {
    case "play":
      play();
      break;
    case "pause":
      pause();
      break;
  }
};

window.addEventListener("load", function(evt) {
  ws_addr = e("ws").value; // read from hidden input field
  addEvent(window, "resize", resizer);
  addEvent(window, "orientationchange", resizer);

  openWebSocket();

  updateProgress();

  // add onclick to every button
  for (var k in views) {
    e(views[k]).onclick = showView(k);
  }
  ["random", "consume", "repeat", "single"].map(function(value) {
    e(value + "Enable").onclick = btnCommand(value, "enable");
    e(value + "Disable").onclick = btnCommand(value, "disable");
  });

  e("submitSearch").onclick = function(evt) {
    return command("search", e("searchText").value);
  };

  e("searchText").onchange = function(evt) {
    return command("search", e("searchText").value);
  };

  ["play", "resume", "pause", "stop", "next", "previous", "clean"].map(function(
    value
  ) {
    e(value).onclick = btnCommand(value, "");
  });

  window.onfocus = function(event) {
    // request a fresh status as some browsers (e. g. Chrome) suspend our
    // progress setTimeout functions
    command("updateRequest", "");
  };

  // show the viewPlaylist
  showId("viewPlaylist");
});

// eof
