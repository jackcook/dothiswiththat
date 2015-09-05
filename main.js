String.prototype.capitalizeFirstLetter = function() {
    return this.charAt(0).toUpperCase() + this.slice(1);
}

var actions = {
  "read a book": "book",
  "watch a movie": "movie",
  "listen to music": "music",
  "read the news": "news",
  "watch a tv show": "tv"
}

var languages = {
  "English": "en",
  "Spanish": "es",
  "French": "fr",
  "Swedish": "sv"
}

$("#actions").autocomplete({
  source: Object.keys(actions)
});

$("#languages").autocomplete({
  source: Object.keys(languages)
});

var go = document.getElementById("go");
var actionsElement = document.getElementById("actions");
var languagesElement = document.getElementById("languages");

go.addEventListener("click", function() {
  var action = actions[actionsElement.value.toLowerCase()];
  var lang = languages[languagesElement.value.toLowerCase().capitalizeFirstLetter()];

  file = "languages/" + lang + ".json";

  request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState == 3) {
      var response = JSON.parse(request.responseText)[action];
      var link = response[Math.floor(Math.random() * response.length)];

      var type = link.split(":")[0];
      var value = link.split(":")[1];
      var url = "nothing.html";

      if (type == "gplaytv") {
        url = "https://play.google.com/store/tv/show?id=" + value;
      } else if (type == "http" || type == "https") {
        url = link;
      } else if (type == "youtube") {
        url = "https://youtube.com/watch?v=" + value;
      }

      window.location = url;
    }
  }

  request.open("GET", file, true);
  request.send();
});
