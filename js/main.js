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
  window.location = "content.html?action=" + actions[actionsElement.value.toLowerCase()] + "&lang=" + languages[languagesElement.value.toLowerCase().capitalizeFirstLetter()];
});
