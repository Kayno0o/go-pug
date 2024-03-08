import http from 'k6/http';
import { check } from 'k6';
import { Rate } from 'k6/metrics';

/** @type {import('k6/options').Options} */
export const defaultOptions = {
  stages: [
    { duration: "5s", target: 15000 },
    { duration: "5s", target: 20000 },
  ],
};

export let options = defaultOptions;

export const errorRate = new Rate('errors');

export default function () {
  const url = 'http://localhost:8080';
  const params = {};
  check(http.get(url, params), {
    'status is 200': (r) => r.status == 200,
  }) || errorRate.add(1);
}
