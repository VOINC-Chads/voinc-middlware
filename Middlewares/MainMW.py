# MainMW
# Purpose is to fulfill requests and message sending.


import zmq


class MainMW():


    def __init__(self):

        self.req = None


    def configure(self):

        try:
            print("Configuring")

        except Exception as e:
            raise e

    def event_loop(self):

        try:

            while True:

                print("Handle events")

        except Exception as e:
            raise e