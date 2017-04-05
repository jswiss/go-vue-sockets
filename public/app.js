new Vue({
  el: '#app',
  data: {
    ws: null, // our websocket
    newMsg: '', // Holds new messages to be sent to the server
    chatContent: '', // A running list of chat messages displayed on the screen
    email: null, // Email address used for grabbing a Gravatar
    username: null, // Our username
    joined: false, // True if email and username have been filled in
  },
  created: function () {
    var self = this;
    this.ws = new Websocket('ws://' + window.location.host + '/ws');
    this.ws.addEventListener('message', function (e) {
      var msg = JSON.parse(e.data);
      self.chatContent += '<div class="chip">' +
        '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
        +
        msg.username +
        '</div>' +
        emojione.toImage(msg.message) + '<br/>'; // Parse emojis
    });
  },
});