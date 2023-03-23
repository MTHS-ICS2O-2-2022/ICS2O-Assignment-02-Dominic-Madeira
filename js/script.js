// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: March 2023
// This file contains the JS functions for index.html


'use strict'
/**
* This function calculates the volume of a cube.
*/
function myButtonClicked () {
 // input
  const edge = parseInt(document.getElementById('side-length').value)

 // process
  const volume = (edge ** 3)

 // output
  document.getElementById('volume').innerHTML = 'The Volume of the Square is: ' + volume + ' cm³'
}
