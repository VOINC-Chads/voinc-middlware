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

    def replica_watching_main(self, obj):

        try:

            print("Need to handle replica watches")

        except Exception as e:
            raise e

    def watch_main_change(self, name, obj):

        try:

            @self.zk.DataWatch(name)
            def data_change(data, stat):

                if data is None:
                    while not self.exists(name):
                        self.logger.info("Waiting for main")
                obj.connect_main(data.decode('utf-8'))



        except Exception as e:
            raise e


    def replica_watch_main(self, name, value, obj):

        try:

            @self.zk.DataWatch(name)
            def main_change(data, stat):

                if data is None:
                    leader = self.create(name, eph=True, value=value)
                    obj.handle_fault(leader)


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


    def get_num_children(self, name):

        try:

            if not self.zk.exists(name):
                return 0
            return len(self.zk.get_children(name))

        except Exception as e:
            raise e


    def get_node_data(self, name):

        try:

            if self.zk.exists(name):
                return self.zk.get(name)
            return None


        except Exception as e:
            raise e





