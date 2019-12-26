'use strict';
const Hostel = require('./hostel');

const URL = 'http://localhost:8080';
const hostel = new Hostel(URL);

// hostel.getBestHostel('biology');
// hostel.getBestHostel('literature');
hostel.sendStudent('Masha', 'computerScience', '3');
