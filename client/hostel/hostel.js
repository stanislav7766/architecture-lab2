const requests = require('../tools/http/http');

class Hostel {
  constructor(URL) {
    this.URL = URL;
  }
  async getBestHostel(speciality) {
    try {
      const res = await requests.getBestHostel(this.URL, speciality);
      console.dir(res, { depth: null });
    } catch (error) {
      console.error(error);
    }
  }
  async sendStudent(name, speciality, hostel) {
    try {
      const res = await requests.sendStudent(
        this.URL,
        name,
        speciality,
        hostel
      );
      console.dir(res, { depth: null });
    } catch (error) {
      console.error(error);
    }
  }
}

module.exports = Hostel;
