#!/usr/bin/env python3

import logging
import argparse
import re
import configparser
import os
import sys
import grpc
from protos.url_broker_pb2 import URL
from protos.url_broker_pb2_grpc import URLBrokerServiceStub

logger = logging.getLogger()
logger.addHandler(logging.StreamHandler(sys.stdout))
logger.setLevel(logging.INFO)

def load_config():
  config = configparser.ConfigParser()
  standard_config_file_locations = [ "/etc/broho/broho.conf",
                                     "~/.config/broho/broho.conf" ]
  for config_file in standard_config_file_locations:
    if os.path.isfile(config_file):
      logger.debug(f"Loaded config file at {config_file}")
      try:
        config.read_file(open(config_file))
        if 'Global' in config.sections():
          return config
      except PermissionError:
        logger.error(f"Couldn't open config file '{config_file}' (permission denied)")
  
  logger.info("No config file found (using defaults)")
  config['Global'] = { 'BindAddress': '127.0.0.1',
                       'BindPort': '50051'
  }
  return config

def main():
  parser = argparse.ArgumentParser()
  parser.add_argument('urls', nargs='+', help='URL(s) to open on host')
  args = parser.parse_args()
  
  config = load_config()
  bind_address = config['Global'].get('BindAddress', '127.0.0.1')
  bind_port = config['Global'].get('BindPort', '50051')
  endpoint = f"{bind_address}:{bind_port}"

  try:
    channel = grpc.insecure_channel(endpoint)
    client = URLBrokerServiceStub(channel)
    for url_string in args.urls:
      if re.match(r'https?://.*', url_string):
        url = URL(url=url_string)
        logger.info(f"Pushing URL {url_string}")
        response = client.PushURL(url)
        if response.response_code != 0:
          logger.error(f"Error occurred when pushing URL {url.url} (response code: {response.response_code}")
      else:
        logger.error(f"Invalid URL: {url_string}")
  except grpc.RpcError as err:
    logger.error(err.details())

# Handle interruption
if __name__ == "__main__":
  try:
    main()
  except KeyboardInterrupt:
    pass
