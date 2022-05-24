#!/usr/bin/env python3
# brohod.py
# Version: 1.0 (2022-05-19)

"""
BRowser On HOst (BROHO): Simple daemon that listens for URL links requests
and launches them in the default browser of the host running the daemon.
"""

__author__ = "Diego Nuevo"
__version__ = "1.0"

import configparser
import webbrowser
import os
import logging
from systemd.journal import JournalHandler
import grpc
from concurrent import futures
from protos import url_broker_pb2_grpc
from protos.url_broker_pb2 import Response
from protos.url_broker_pb2_grpc import URLBrokerServiceServicer

# Set up global logging
SERVICE_NAME = 'brohod'
logger = logging.getLogger(SERVICE_NAME)
logger.addHandler(JournalHandler(SYSLOG_IDENTIFIER=SERVICE_NAME))
logger.setLevel(logging.INFO)

class URLBrokerService(URLBrokerServiceServicer):
  def PushURL(self, request, context):
    if request.url.startswith('http'):
      logger.debug(f"Opening URL '{request.url}' in a new tab using the default browser")
      webbrowser.open(request.url, new=2)
      return Response(response_code=0)
    else:
      logger.error(f'Invalid URL: {request.url}')
      return Response(response_code=1)

def load_config():
  config = configparser.ConfigParser()
  standard_config_file_locations = [ "/etc/broho/broho.conf",
                                     "~/.config/broho/broho.conf" ]
  for config_file in standard_config_file_locations:
    if os.path.isfile(config_file):
      logger.info(f"Loaded config file at {config_file}")
      try:
        config.read_file(open(config_file))
        return config
      except PermissionError:
        logger.error(f"Couldn't open config file '{config_file}' (permission denied)")
  
  #No config file found, using default values
  logger.warning("No config file found (using defaults)")
  config['Global'] = { 'BindAddress': '127.0.0.1',
                       'BindPort': '50051'
  }
  return config

def run_server():
  config = load_config()
  endpoint = f"{config['Global']['BindAddress']}:{config['Global']['BindPort']}"

  server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
  url_broker_pb2_grpc.add_URLBrokerServiceServicer_to_server(URLBrokerService(), server)
  server.add_insecure_port(endpoint)
  logger.info(f"Starting {SERVICE_NAME} server on {endpoint}...")
  server.start()
  server.wait_for_termination()
  logger.info(f"Stopping {SERVICE_NAME} server on {endpoint}...")

# Handle interruption
if __name__ == "__main__":
  try:
    run_server()
  except KeyboardInterrupt:
    pass