var timerElement = document.getElementById('timer');
var startButton = document.getElementById('start-button');
var stopButton = document.getElementById('stop-button');

var timer;
var seconds = 0;
var minutes = 25;
var isRunning = false;

startButton.addEventListener('click', function() {
    if (!isRunning) {
        isRunning = true;
        timer = setInterval(updateTimer, 1000);
    }
});

stopButton.addEventListener('click', function() {
    clearInterval(timer);
    isRunning = false;
    timerElement.innerHTML = '25:00';
    seconds = 0;
    minutes = 25;
});

function updateTimer() {
    seconds--;

    if (seconds < 0) {
        seconds = 59;
        minutes--;
    }

    if (minutes < 0) {
        clearInterval(timer);
        isRunning = false;
        timerElement.innerHTML = '00:00';
        alert('Time is up!');
        return;
    }

    var formattedSeconds = seconds < 10 ? '0' + seconds : seconds;
    var formattedMinutes = minutes < 10 ? '0' + minutes : minutes;

    timerElement.innerHTML = formattedMinutes + ':' + formattedSeconds;
}
