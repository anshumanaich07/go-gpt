import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
// import {Prompt} from '../../wailsjs/go/main/App';
import {Prompt} from '../wailsjs/go/ui/App'

// TODO: add a login

document.querySelector('#app').innerHTML = `
  <img id="logo" class="logo">
    <div class="input-box" id="input" >
      <input class="input" id="prompt" type="text" autocomplete="off" placeholder="Enter your query" />
      <button class="btn" onclick="prompt()">Go</button>
    </div>
    <br />
    <div class="result-grp" id="result-grp">
      <div id="result-prompt"></div>
    </div>
  </div>
`;
document.getElementById('logo').src = logo;

// GET INPUT VALUES
let promptEle = document.getElementById("prompt")
let resultElement = document.getElementById("result-prompt");

window.prompt = function () {
    let prompt = promptEle.value;
    if (prompt === "") {
      // TODO: set up the warning in div
      return "Please enter a prompt!"
    }
    try {
      Prompt(prompt)
        .then((prompt) => {
          resultElement.innerText = prompt;
        })
        .catch((err) => {
          console.error(err);
        });
    } catch (err) {
        console.error(err);
    }
};