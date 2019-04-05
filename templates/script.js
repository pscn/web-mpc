window.addEventListener('load', function (evt) {
  // short for document.getElementById
  var e = function (id) {
    return document.getElementById(id)
  }

  var ws_addr = e('ws').value // read from hidden input field
  var ws


  // https://www.w3schools.com/js/js_cookies.asp
  var getCookie = function (cookieName) {
    var name = cookieName + '='
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
      e('configHost').value = parts[0]
      e('configPort').value = parts[1]
      e('configPass').value = parts[2]
    }
  }
  updateConfig()

  var command = function (cmd, data) {
    if (!ws) { return false; }
    myJson = { 'command': cmd, 'data': data }
    // console.log('SEND: ' + JSON.stringify(myJson))
    ws.send(JSON.stringify(myJson))
    return false
  }

  var previousDirectory = '' // for browsing

  // a few 'globals' to track the process and current state of the player
  var gDuration = 1.0
  var gElapsed = 0.0
  var gState = 'pause'
  var readableSeconds = function (value) {
    var min = parseInt(value / 60)
    var sec = parseInt(value % 60)
    if (sec < 10) { sec = '0' + sec; }
    return min + ':' + sec
  }
  var updateProgress = function () {
    e('elapsed').innerHTML = readableSeconds(gElapsed)
    e('duration').innerHTML = readableSeconds(gDuration)
    if ((gState == 'play') && (gElapsed < gDuration)) {
      gElapsed += 1.0
    }
    setTimeout(updateProgress, 1000)
  }

  e('connectionStatus').onclick = function () {
    openWebSocket()
  }
  var showError = function (msg) {
    e('error').innerHTML = e('error').innerHTML + '<br />' + msg
    e('error').style.display = ''
  }
  var hideError = function () {
    e('error').innerHTML = ''
    e('error').style.display = 'none'
  }

  var updateStatus = function (data) {
    const { state, duration, elapsed, consume, repeat, random, single } = data
    gState = state
    console.log(`updateStatus(${state})`)
    switch (state) {
      case 'pause':
      case 'play':
        togglePlayPause(state)
        gDuration = duration
        gElapsed = elapsed
        break
      case 'stop':
        stop()
        gDuration = 0.0
        gElapsed = 0.0
        break
    }
    // update the config view
    e('consume').checked = consume ? 'checked' : '';
    e('repeat').checked = repeat ? 'checked' : '';
    e('random').checked = random ? 'checked' : '';
    e('single').checked = single ? 'checked' : '';
  }
  var updateActiveSong = function (data) {
    const { file, artist, title, album_artist, album } = data
    e('activeSong').title = file
    e('artist').innerHTML = artist
    e('title').innerHTML = title
    if (artist != album_artist) { // only show album artist if it's different
      e('albumArtist').style.display = ''
      e('albumArtist').innerHTML = '[' + album_artist + ']&nbsp;'
    } else {
      e('albumArtist').style.display = 'none'
    }
    e('album').innerHTML = album
  }

  var newSongNode = function (id, entry) {
    const { file, artist, title, album, album_artist, duration } = entry;
    var node = e(id).cloneNode(true)
    node.style.display = ''
    node.title = file
    node.querySelector('#songCellArtist').innerHTML = artist
    node.querySelector('#songCellTitle').innerHTML = title
    node.querySelector('#songCellAlbum').innerHTML = album
    if (artist != album_artist) { // only show album artist if it's different
      node.querySelector('#songCellAlbumArtist').innerHTML =
        '[' + album_artist + ']&nbsp;'
    } else {
      node.querySelector('#songCellAlbumArtist').style.display = 'none'
    }
    node.querySelector('#songCellArtist').innerHTML = artist
    node.querySelector('#songCellDuration').innerHTML =
      readableSeconds(duration)
    return node
  }
  var btnCommand = function (cmd, data) {
    return function () {
      return command(cmd, data)
    }
  }

  var processResponse = function (obj) {
    console.log({ obj })
    const { type, data } = obj

    switch (type) {
      case 'error', 'info':
        showError(data)
        break
      case 'status':
        updateStatus(data)
        break
      case 'activeSong':
        updateActiveSong(data)
        break
      case 'activePlaylist':
        e('playlist').innerHTML = '' // delete old playlist
        data.Playlist.map(function (entry, i) {
          const { file } = entry
          var node = newSongNode('playlistEntry', entry)
          // disable the play button for the active song
          node.querySelector('#plPlay').disabled = (file == e('activeSong').title)
            ? 'disabled' : ''
          node.querySelector('#plPlay').onclick = btnCommand('play', i.toString())
          node.querySelector('#plRemove').onclick = btnCommand('remove', i.toString())
          e('playlist').append(node)
        })
        break
      case 'searchResult':
        e('searchResult').innerHTML = '' // delete old search result
        data.SearchResult.map(function (entry) {
          const { file } = entry
          var node = newSongNode('searchEntry', entry)
          node.querySelector('#srAdd').onclick = btnCommand('add', file)
          e('searchResult').append(node)
        })
        break
      case 'directoryList':
        e('directoryList').innerHTML = ''
        if (previousDirectory != '') {
          node = e('directoryListEntry').cloneNode(true)
          node.id = 'dlRowParent'
          node.style.display = ''
          node.querySelector('#dlName').innerHTML = previousDirectory
          node.querySelector('#dlBrowse').onclick = btnCommand('browse', name)
          e('directoryList').append(node)
        }
        data.directoryList.map(function (entry, i) {
          var node
          if (entry.type == 'directory') {
            node = e('directoryListEntry').cloneNode(true)
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
            node = e('searchEntry').cloneNode(true)
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
          e('directoryList').append(node)
        })
        break
    }
  }
  var openWebSocket = function () {
    ws = new WebSocket(ws_addr)
    ws.onopen = function (evt) {
      console.log('OPEN')
      hideError()
      e('connect').disabled = 'disabled'
    }
    ws.onclose = function (evt) {
      console.log('CLOSE')
      ws = null
      showError('no connection')
      e('connect').disabled = ''
    }
    ws.onmessage = function (evt) {
      processResponse(JSON.parse(evt.data))
    }
    ws.onerror = function (evt) {
      showError(evt.data)
    }
  }
  openWebSocket()

  window.onfocus = function (event) {
    // request a fresh status as some browsers (e. g. Chrome) suspend our
    // progress setTimeout functions
    command('statusRequest', '')
  }
  updateProgress()

  // add onclick function for all controls
  var togglePlayPause = function (state) {
    console.log(`togglePlayPause(${state})`)
    switch (state) {
      case 'play':
        e('pause').style.display = ''
        e('resume').style.display = 'none'
        break
      case 'pause':
        e('pause').style.display = 'none'
        e('resume').style.display = ''
        break
    }
    e('play').style.display = 'none'
    e('stop').disabled = ''
    e('next').disabled = ''
    e('previous').disabled = ''
  }

  var stop = function () {
    e('play').style.display = ''
    e('pause').style.display = 'none'
    e('resume').style.display = 'none'
    e('stop').disabled = 'disabled'
    e('next').disabled = 'disabled'
    e('previous').disabled = 'disabled'
  }

  var buttonIDs = ['play', 'resume', 'pause', 'stop', 'next', 'previous']
  buttonIDs.map(function (value) {
    document.getElementById(value).onclick = function (evt) {
      console.log(`Control: ${value}`)
      return command(value, '')
    }
  })

  var views = {
    'playlistView': 'list',
    'searchView': 'search',
    'folderView': 'browser',
    'configView': 'config'
  }
  var show = function (view) {
    return function () {
      switch (view) {
        case 'folderView':
          command('browse', '')
          break
      }
      for (var k in views) {
        switch (k) {
          case view: // show matching view
            e(k).style.display = ''
            e(views[k]).disabled = 'disabled'
            break
          default: // hidde others
            e(k).style.display = 'none'
            e(views[k]).disabled = ''
            break
        }

      }
    }
  }

  e('list').onclick = show('playlistView')
  e('search').onclick = show('searchView')
  e('browser').onclick = show('folderView')
  e('config').onclick = show('configView')

  e('submitSearch').onclick = function (evt) {
    return command('search', e('searchText').value)
  }

  e('submitConfig').onclick = function (evt) {
    // update the cookie and reconnect to the websocket
    // which will read the cookie and so apply the changes
    document.cookie = 'mpd=' + e('configHost').value + ':'
      + e('configPort').value + ':'
      + e('configPass').value
    if (ws != null) {
      ws.close()
    }
    setTimeout(openWebSocket, 1000)
  }
  e('searchText').onchange = function (evt) {
    return command('search', e('searchText').value)
  }

  show('playlistView')
})

// eof
