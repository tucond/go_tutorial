//import { initializeApp } from "firebase/compat/app";
import firebase from "firebase/compat/app";
import "firebase/compat/auth"

const firebaseConfig = {
    apiKey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    authDomain: "xxxxxxx-xxxxxx.firebaseapp.com",
    databaseURL: "https://xxxxxxx-xxxxxx.firebaseio.com",
    projectId: "xxxxxxx-xxxxxxx-xxxxxx",
    storageBucket: "",
    messagingSenderId: "xxxxxxxxxxxx",
    appId: "1:xxxxxxxxxxxx:web:xxxxxxxxxxxxxxxx""
  };
  
  // Initialize Firebase
  firebase.initializeApp(firebaseConfig);

export default firebase;