const fetch = require('node-fetch');

module.exports = {
  getBestHostel: (URL, speciality) =>
    new Promise((resolve, reject) => {
      fetch(`${URL}/getBestHostel`, {
        method: 'POST',
        body: JSON.stringify({ speciality }),
        headers: { 'Content-Type': 'application/json' }
      })
        .then(res => resolve(res.json()))
        .catch(err => reject(err));
    }),
  sendStudent: (URL, name, speciality, hostel) =>
    new Promise((resolve, reject) =>
      fetch(`${URL}/sendStudent`, {
        method: 'POST',
        body: JSON.stringify({ name, speciality, hostel }),
        headers: { 'Content-Type': 'application/json' }
      })
        .then(res => resolve(res.json()))
        .catch(err => reject(err))
    )
};
