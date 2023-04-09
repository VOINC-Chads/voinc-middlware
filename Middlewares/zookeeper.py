# zookeeper
# Zookeeper adapter logic
import kazoo.exceptions
from kazoo.client import KazooClient

class ZK():

    def __init__(self, port, addr, logger):

        self.logger = logger
        self.zk = None
        self.port = port
        self.addr = addr

        hosts = self.addr + ":" + str(self.port)

        self.zk = KazooClient(hosts)

        self.zk.start()

    def watch_main_change(self, obj):

        try:

            @self.zk.DataWatch("/main")
            def data_change(data, stat):

                if data is None:
                    while not self.exists("/main"):
                        self.logger.info("Waiting for main")
                obj.connect_main(data)



        except Exception as e:
            raise e


    def create(self, name, eph, value):

        try:
            if self.exists(name):
                return False
            else:
                self.zk.create(path=name, ephemeral=eph, value=bytes(value, 'utf-8'), makepath=True)
                return True

        except kazoo.exceptions.NodeExistsError:
            return False
        except Exception as e:
            raise e
    def exists(self, name):

        try:

            if (self.zk.exists(name)):
                return True
            return False

        except Exception as e:
            raise e





