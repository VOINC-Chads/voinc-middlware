# Main.cpp
# Main/master of the volunteer computing architecture
# Volunteers provide registration information and then main distributes jobs from workers
# to volunteers. Due to bad actors, it is the responsibility of a consensus of responses to be received
# to be sent back to the worker.

import argparse
import zmq
from Middleware.MainMW import MainMW
from Middleware.zookeeper import ZK
import logging
from Messages import messages_pb2

class Main():


    def __init__(self, logger):

        self.logger = logger
        self.volunteers = {}
        self.pending = {}
        self.mw_obj = None
        self.name = None

        self.numVolunteers = None
        self.numOccupied = None

        self.zkAddr = None
        self.zkPort = None
        self.zk = None

        self.quorum = None

    def set_pending(self, id):

        try:

            self.logger.info("Main::set_pending - setting job came in from id mentioned {}".format(id))

            if id not in self.pending:
                self.pending[id] = {}

        except Exception as e:
            raise e
    def quorum_met(self, main_resp):

        try:
            self.logger.info("Main::quorum_met - checking if value has quorum met")
            resp = messages_pb2.MainResp()
            resp.ParseFromString(main_resp[1])
            id = main_resp[0]

            job_resp = resp.job_resp

            self.logger.info(self.pending)

            for result in job_resp.results:

                value = result.value
                result = result.result

                self.logger.info("Checking quorum for {}".format(value))

                if value not in self.pending[id]:
                    self.pending[id][value] = []
                self.pending[id][value].append(result)



                self.logger.info("Values present after append: {}".format(self.pending[id][value]))

                if len(self.pending[id][value]) == self.quorum:
                    self.logger.info("Quorum size met")
                    consensus_result = self.find_consensus(value, id)
                    self.pending[id].pop(value)
                    self.mw_obj.send_job_resp(main_resp[0], value, consensus_result)
                    self.logger.info(self.pending)
                    return True

            self.logger.info("Main::quorum_met - quorum not met, so nothing sent back")
            return False




        except Exception as e:
            raise e


    def get_heartbeat_info(self):
        try:

            jobs = 0
            for id in self.pending:
                jobs += len(self.pending[id].keys())

            workers = len(self.volunteers.keys())

            return jobs, workers

        except Exception as e:
            raise e
    def find_consensus(self, value, id):

        try:

            results = self.pending[id][value]
            occurrences = {}
            for val in results:
                if val not in occurrences:
                    occurrences[val] = 0
                occurrences[val] += 1

            max_result = max(occurrences, key=occurrences.get)

            return max_result


        except Exception as e:
            raise e
    def configure(self, args):

        try:

            self.logger.info("Main::configure - configuring main")

            self.numOccupied = 0
            self.numVolunteers = 0
            self.name = args.name
            self.quorum = args.quorum

            self.zkAddr = args.zkaddr
            self.zkPort = args.zkport

            self.zk = ZK(self.zkPort, self.zkAddr, self.logger)

            self.mw_obj = MainMW(self.logger)
            self.mw_obj.configure(args)



        except Exception as e:
            raise e


    def register_volunteer(self, information, id):

        try:

            self.logger.info("Main::register_volunteer - registering a volunteer below")
            self.logger.info(information)

            name = information.info.id
            port = information.info.port
            addr = information.info.addr
            capacity = information.info.capacity

            successful = True

            if str(addr) + ":" + str(port) in self.volunteers:
                self.logger.info("Already registered worker")
                successful = False
            else:
                information = [name, port, addr, capacity]
                self.volunteers[str(addr) + ":" + str(port)] = information
                self.mw_obj.set_dealer(str(addr) + ":" + str(port))

            self.logger.info(self.volunteers)

            self.mw_obj.send_register_response(successful, id)


        except Exception as e:
            raise e


    def driver(self):

        try:

            self.logger.info("Main::driver - entered driver")


            # Set the upcall handler from the middleware
            self.mw_obj.set_upcall_handle(self)

            # start the event loop that just continuously runs and checks for requests

            self.mw_obj.event_loop(timeout=None)



        except Exception as e:
            raise e

def parseCmdLineArgs():


    parser = argparse.ArgumentParser(description="Main/Master Application")

    parser.add_argument("-n", "--name", default="main", help="Name assigned to master")
    parser.add_argument("-a", "--addr", default="localhost", help="Address process is running on")
    parser.add_argument("-p", "--port", default="8000", help="Port process is running on")
    parser.add_argument("-z", "--zkaddr", default="localhost", help="Address zookeeper is running on")
    parser.add_argument("-o", "--zkport", default="2181", help="Port zookeeper is running on")
    parser.add_argument("-l", "--loglevel", type=int, default=logging.INFO,
                        choices=[logging.DEBUG, logging.INFO, logging.WARNING, logging.ERROR, logging.CRITICAL],
                        help="logging level, choices 10,20,30,40,50: default 20=logging.INFO")
    parser.add_argument("-q", "--quorum", type=int, default=1, help="quorum size")
    parser.add_argument("-s", "--leadersize", type=int, default=1, help="Number of replicas needed to begin")

    return parser.parse_args()


def main():
    try:
        # obtain a system wide logger and initialize it to debug level to begin with
        logging.debug("Main - acquire a child logger and then log messages in the child")
        logger = logging.getLogger("Main")

        # first parse the arguments
        logger.debug("Main: parse command line arguments")
        args = parseCmdLineArgs()

        # reset the log level to as specified
        logger.debug("Main: resetting log level to {}".format(args.loglevel))
        logger.setLevel(args.loglevel)
        logger.debug("Main: effective log level is {}".format(logger.getEffectiveLevel()))

        # Obtain a publisher application
        logger.debug("Main: obtain the object")
        disc_app = Main(logger)

        # configure the object
        disc_app.configure(args)

        # now invoke the driver program
        disc_app.driver()

    except Exception as e:
        logger.error("Exception caught in main - {}".format(e))
        return

if __name__ == "__main__":

  # set underlying default logging capabilities
  logging.basicConfig (level=logging.DEBUG,
                       format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')


  main()
