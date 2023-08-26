import { sleep } from 'k6';
import http from 'k6/http';

export default function () {
  //List all urls
  http.get('http://localhost:8080/url-short');


  let body = {
    "url": "www.facebook.com.br",
    "expiration": "2023-12-31 23:59:59"
  }
  let res = http.post('http://localhost:8080/url-short', JSON.stringify(body))

  
  sleep(1)
}

// let data = { name: 'Bert' };

// // Using a JSON string as body
// let res = http.post(url, JSON.stringify(data), {
//   headers: { 'Content-Type': 'application/json' },
// });