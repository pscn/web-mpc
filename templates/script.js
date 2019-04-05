window.addEventListener('load', function (evt) {
  var ws

  // Functions
  // short for document.getElementById
  var el = function (id) {
    return document.getElementById(id)
  }

  // https://www.w3schools.com/js/js_cookies.asp
  var getCookie = function (cname) {
    var name = cname + '='
    var decodedCookie = decodeURIComponent(document.cookie)
    var ca = decodedCookie.split(';')
    for (var i = 0; i < ca.length; i++) {
      var c = ca[i]
      while (c.charAt(0) == ' ') {
        c = c.substring(1)
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length)
      }
    }
    return ''
  }

  var updateConfig = function () {
    mpdAddr = getCookie('mpd')
    if (mpdAddr != '') {
      parts = mpdAddr.split(':')
      el('configHost').value = parts[0]
      el('configPort').value = parts[1]
      el('configPass').value = parts[2]
    }
  }
  updateConfig()

  var command = function (cmd, data) {
    if (!ws) { return false; }
    myJson = { 'command': cmd, 'data': data }
    console.log('SEND: ' + JSON.stringify(myJson))
    ws.send(JSON.stringify(myJson))
    return false
  }

  var previousDirectory = '' // for browsing

  // a few "globals" to track the process and current state of the player
  var duration = 1.0
  var elapsed = 0.0
  var state = 'pause'
  var readableSeconds = function (value) {
    var min = parseInt(value / 60)
    var sec = parseInt(value % 60)
    if (sec < 10) { sec = '0' + sec; }
    return min + ':' + sec
  }
  var updateProgress = function () {
    el('elapsed').innerHTML = readableSeconds(elapsed)
    el('duration').innerHTML = readableSeconds(duration)
    if ((state == 'play') && (elapsed < duration)) {
      elapsed += 1.0
    }
    setTimeout(updateProgress, 1000)
  }

  el('connectionStatus').onclick = function () {
    openWebSocket()
  }
  var showError = function (msg) {
    el('error').innerHTML = el('error').innerHTML + '<br />' + msg
    el('error').style.display = ''
  }
  var hideError = function () {
    el('error').innerHTML = ''
    el('error').style.display = 'none'
  }

  var updateActiveSong = function (data) {
    el('activeSong').title = data.file
    el('artist').innerHTML = data.artist
    el('title').innerHTML = data.title
    if (obj.data.artist != data.album_artist) {
      el('albumArtist').style.display = ''
      el('albumArtist').innerHTML = '[' + data.album_artist + ']&nbsp;'
    } else {
      el('albumArtist').style.display = 'none'
    }
    el('album').innerHTML = obj.data.album
  }

  var newSongNode = function (id, entry) {
    var node = el(id).cloneNode(true)
    node.style.display = ''
    node.querySelector('#songCellArtist').innerHTML = entry.artist
    node.querySelector('#songCellTitle').innerHTML = entry.title
    node.querySelector('#songCellAlbum').innerHTML = entry.album
    if (entry.artist != entry.album_artist) {
      node.querySelector('#songCellAlbumArtist').innerHTML =
        '[' + entry.album_artist + ']&nbsp;'
    } else {
      node.querySelector('#songCellAlbumArtist').style.display = 'none'
    }
    node.querySelector('#songCellArtist').innerHTML = entry.artist
    node.querySelector('#songCellDuration').innerHTML =
      readableSeconds(entry.duration)
    return node
  }
  var openWebSocket = function () {
    var ws_addr = el('ws').value
    ws = new WebSocket(ws_addr)
    ws.onopen = function (evt) {
      console.log('OPEN')
      hideError()
      el('connect').disabled = 'disabled'
    }
    ws.onclose = function (evt) {
      console.log('CLOSE')
      ws = null
      showError('no connection')
      el('connect').disabled = ''
    }
    ws.onmessage = function (evt) {
      console.log('RESPONSE: ' + evt.data)
      obj = JSON.parse(evt.data)
      switch (obj.type) {
        case 'error', 'info':
          showError(obj.data)
          break
        case 'status':
          switch (obj.data.state) {
            case 'pause':
              pause()
              state = 'pause'
              duration = obj.data.duration
              elapsed = obj.data.elapsed
              break
            case 'play':
              play()
              state = 'play'
              duration = obj.data.duration
              elapsed = obj.data.elapsed
              break
            case 'stop':
              stop()
              state = 'stop'
              duration = 0.0
              elapsed = 0.0
              break
          }
          if (obj.data.consume) {
            el('consume').checked = 'checked'
          } else {
            el('consume').checked = ''
          }
          if (obj.data.repeat) {
            el('repeat').checked = 'checked'
          } else {
            el('repeat').checked = ''
          }
          if (obj.data.random) {
            el('random').checked = 'checked'
          } else {
            el('random').checked = ''
          }
          if (obj.data.single) {
            el('single').checked = 'checked'
          } else {
            el('single').checked = ''
          }
          break
        case 'activeSong':
          updateActiveSong(obj.data)
          break
        case 'activePlaylist':
          el('playlist').innerHTML = ''
          obj.data.Playlist.map(function (entry, i) {
            var node = newSongNode('playlistEntry', entry)
            {
              const j = i
              const file = entry.file
              if (file == el('activeSong').title) {
                node.querySelector('#plPlay').disabled = 'disabled'
              } else {
                node.querySelector('#plPlay').disabled = ''
              }
              node.querySelector('#plPlay').onclick = function (evt) {
                return command('play', j.toString())
              }
              node.querySelector('#plRemove').onclick = function (evt) {
                return command('remove', j.toString())
              }
            }
            el('playlist').append(node)
          })
          break
        case 'searchResult':
          el('searchResult').innerHTML = ''
          obj.data.SearchResult.map(function (entry, i) {
            var node = newSongNode('searchEntry', entry)
            {
              const file = entry.file
              node.querySelector('#srAdd').onclick = function (evt) {
                return command('add', file)
              }
            }
            el('searchResult').append(node)
          })
          break
        case 'directoryList':
          el('directoryList').innerHTML = ''
          if (previousDirectory != '') {
            node = el('directoryListEntry').cloneNode(true)
            node.id = 'dlRowParent'
            node.style.display = ''
            node.querySelector('#dlName').innerHTML = previousDirectory
            {
              const name = parent
              node.querySelector('#dlBrowse').onclick = function (evt) {
                return command('browse', name)
              }
            }
            el('directoryList').append(node)
          }
          obj.data.directoryList.map(function (entry, i) {
            var node
            if (entry.type == 'directory') {
              node = el('directoryListEntry').cloneNode(true)
              node.id = 'dlRow' + i
              node.style.display = ''
              node.querySelector('#dlName').innerHTML = entry.directory

              {
                const name = entry.directory
                node.querySelector('#dlBrowse').onclick = function (evt) {
                  return command('browse', name)
                }
              }
            } else {
              node = el('searchEntry').cloneNode(true)
              node.id = 'srRow' + i
              node.style.display = ''
              node.querySelector('#srArtist').innerHTML = entry.artist
              node.querySelector('#srTitle').innerHTML = entry.title
              node.querySelector('#srAlbum').innerHTML = entry.album
              if (entry.artist != entry.album_artist) {
                node.querySelector('#srAlbumArtist').innerHTML = '[' + entry.album_artist + ']&nbsp;'
              } else {
                node.querySelector('#srAlbumArtist').style.display = 'none'
              }
              node.querySelector('#srArtist').innerHTML = entry.artist
              node.querySelector('#srDuration').innerHTML = readableSeconds(entry.duration)
              {
                const file = entry.file
                node.querySelector('#srAdd').onclick = function (evt) {
                  return command('add', file)
                }
              }
            }
            el('directoryList').append(node)
          })
          break
      }
    }
    ws.onerror = function (evt) {
      showError(evt.data)
    }
  }
  openWebSocket()

  window.onfocus = function (event) {
    // request a fresh status as some browsers (e. g. Chrome) suspend our
    // progress bar setTimeout functions
    command('statusRequest', '')
  }
  updateProgress()

  // add onclick function for all controls
  var pause = function () {
    el('play').style.display = 'none'
    el('pause').style.display = 'none'
    el('resume').style.display = ''
    el('stop').style.display = ''; el('stop').disabled = ''
    el('next').style.display = ''; el('next').disabled = ''
    el('previous').style.display = ''; el('previous').disabled = ''
  }

  var play = function () {
    el('play').style.display = 'none'
    el('pause').style.display = ''
    el('resume').style.display = 'none'
    el('stop').style.display = ''; el('stop').disabled = ''
    el('next').style.display = ''; el('next').disabled = ''
    el('previous').style.display = ''; el('previous').disabled = ''
  }

  var stop = function () {
    el('play').style.display = ''
    el('pause').style.display = 'none'
    el('resume').style.display = 'none'
    el('stop').style.display = ''; el('stop').disabled = 'disabled'
    el('next').style.display = ''; el('next').disabled = 'disabled'
    el('previous').style.display = ''; el('previous').disabled = 'disabled'
  }

  var buttonIDs = ['play', 'resume', 'pause', 'stop', 'next', 'previous']
  buttonIDs.map(function (value) {
    document.getElementById(value).onclick = function (evt) {
      console.log(value)
      return command(value, '')
    }
  })

  var showList = function (evt) {
    el('playlistView').style.display = ''
    el('searchView').style.display = 'none'
    el('folderView').style.display = 'none'
    el('configView').style.display = 'none'

    // enable/disable buttons
    el('search').disabled = ''
    el('list').disabled = 'disabled'
    el('browser').disabled = ''
    el('config').disabled = ''
  }

  var showSearch = function (evt) {
    el('playlistView').style.display = 'none'
    el('searchView').style.display = ''
    el('folderView').style.display = 'none'
    el('configView').style.display = 'none'

    el('searchText').focus()
    el('searchText').select()
    el('searchResult').innerHTML = ''
    el('searchResult').style.display = ''
    // enable/disable buttons
    el('search').disabled = 'disabled'
    el('list').disabled = ''
    el('browser').disabled = ''
    el('config').disabled = ''
  }
  var showBrowser = function (evt) {
    previousDirectory = ''
    el('playlistView').style.display = 'none'
    el('searchView').style.display = 'none'
    el('folderView').style.display = ''
    el('configView').style.display = 'none'

    // enable/disable buttons // enable/disable buttons
    el('search').disabled = ''
    el('list').disabled = ''
    el('browser').disabled = 'disabled'
    el('config').disabled = ''

    command('browse', '')
  }
  var showConfig = function (evt) {
    el('playlistView').style.display = 'none'
    el('searchView').style.display = 'none'
    el('folderView').style.display = 'none'
    el('configView').style.display = ''

    // enable/disable buttons // enable/disable buttons
    el('search').disabled = ''
    el('list').disabled = ''
    el('browser').disabled = ''
    el('config').disabled = 'disabled'
  }

  el('search').onclick = showSearch

  el('browser').onclick = showBrowser

  el('config').onclick = showConfig

  el('list').onclick = showList

  el('submitSearch').onclick = function (evt) {
    return command('search', el('searchText').value)
  }
  el('submitConfig').onclick = function (evt) {
    document.cookie = 'mpd=' + el('configHost').value + ':'
      + el('configPort').value + ':'
      + el('configPass').value
    if (ws != null) {
      ws.close()
    }
    setTimeout(openWebSocket, 1000)
  }
  el('searchText').onchange = function (evt) {
    return command('search', el('searchText').value)
  }

  showList()
})

// eof
