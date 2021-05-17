import logging
import os
from redis import Redis
import requests
import time

DEBUG = os.environ.get("DEBUG", "").lower().startswith("y")

log = logging.getLogger(__name__)
if DEBUG:
    logging.basicConfig(level=logging.DEBUG)
else:
    logging.basicConfig(level=logging.INFO)
    logging.getLogger("requests").setLevel(logging.WARNING)

redis = Redis("redis")


def get_coordinates():
    r = requests.get("http://coordinates:3000/getPoint")
    print('r.content => ')
    print(r.content)
    return r.content


def post_encrypt_coordinates(coordinates):
    r = requests.post("http://encrypter:8080/",
                      data=coordinates,
                      headers={"Content-Type": "application/json"})
    encrypt_coordinates = r.text
    print('encrypt_coordinates => ')
    print(encrypt_coordinates)
    return encrypt_coordinates


def work_loop(interval=1):
    deadline = 0
    loops_done = 0
    while True:
        if time.time() > deadline:
            log.info("{} units of work done, updating coordinates counter"
                     .format(loops_done))
            redis.incrby("encrypt_coordinates", loops_done)
            loops_done = 0
            deadline = time.time() + interval
        work_once()
        loops_done += 1


def work_once():
    log.debug("Doing one unit of work")
    time.sleep(0.1)
    coordinates = get_coordinates()
    hex_encrypt_coordinates = post_encrypt_coordinates(coordinates)
    log.info("Coordinate found: {}...".format(hex_encrypt_coordinates))
    created = redis.hset("points", hex_encrypt_coordinates, coordinates)
    print('redis get start => ')
    print(redis.hget("points", hex_encrypt_coordinates))
    print('redis get end => ')
    if not created:
        log.info("we already registered this coordinate before")


if __name__ == "__main__":
    while True:
        try:
            work_loop()
        except:
            log.exception("In work loop:")
            log.error("Waiting 10s and restarting.")
            time.sleep(10)
