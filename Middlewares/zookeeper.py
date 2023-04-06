# zookeeper
# Zookeeper adapter logic

from kazoo.client import KazooClient

class ZK():

    def __init__(self, port, addr):

        self.zk = None
        self.port = None
        self.addr = None

        
