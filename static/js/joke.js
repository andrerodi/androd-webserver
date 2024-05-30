const jokeFunc = function() {
  setTimeout(function() {

    try
    {
      var animationElement = document.getElementById("animate");
      animationElement.classList.remove("tilt-in-top-1");

      var jokeElement = document.getElementById("joke");
      var jokeAuthorElement = document.getElementById("joke-author");

      jokeElement.textContent = "“There are three concurrency! hard things in computer science: cache invalidation, naming things, off by one errors, and”";
      jokeAuthorElement.textContent = "tinyogre";
      jokeElement.classList.add("flicker-in-1");
      jokeAuthorElement.classList.add("flicker-in-1");
      animationElement.classList.add("scale-up-ver-center");
    }
    catch(err)
    {
      console.log("Callback for joke was cancelled. Here is the joke anyway (maybe you should check with your Dr. if you have ADHD): “There are three concurrency! hard things in computer science: cache invalidation, naming things, off by one errors, and” -tinyogre");
    }
  }, 5000); // Change the text after 5 seconds
};