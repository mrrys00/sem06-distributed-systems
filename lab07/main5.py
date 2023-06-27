import subprocess
import os
from threading import Thread
from time import sleep
from kazoo.client import KazooClient
from kazoo.recipe.watchers import ChildrenWatch

GRAPHICAL_APP_PATH = "/usr/bin/gedit"

class CustomThread(Thread):
    def __init__(self):
        Thread.__init__(self)
        self.value = None

    def run(self):
        sleep(1)
        self.value = subprocess.Popen([GRAPHICAL_APP_PATH, "&"], shell=True)

threads_ = []

def start_external_app():
    try:
        t = CustomThread()
        t.start()
        t.join()
        threads_.append(t.value)
    except:
        pass

def stop_external_app():
    try:
        print(threads_)
        t = threads_[len(threads_)-1]
        print(f"process to be terminated {t}")
        t.terminate()
    except:
        print("no nie zamyka się")


def display_num_children(zk):
    last_clihd_num = 0
    while True:
        num_children = -1
        try:
            num_children = len(zk.get_children("/z/"))
        except:
            pass
        if num_children > last_clihd_num:
            print("Aktualna ilość potomków: {}".format(num_children))
            os.system(f"zenity --info --text={num_children}")
        
        last_clihd_num = num_children


def display_tree_structure(zk, path="/"):
    children = zk.get_children(path)
    print("{}{}".format(path, children))
    for child in children:
        child_path = path + "/" + child
        display_tree_structure(zk, child_path)


def watch_znode(event):
    if event.type == "CREATED":
        start_external_app()
        print("started app")
    elif event.type == "DELETED":
        stop_external_app()
        print("closed app")


def watch_children(event):
    display_num_children(event)


def main():
    zk = KazooClient(hosts='127.0.0.1:2181')
    zk.start()

    display_tree_structure(zk)

    # ChildrenWatch(client=zk, path="/z", func=watch_children)
    t1=Thread(target=display_num_children, args=(zk,))
    t1.start()
    try:
        while True:
            zk.exists("/z", watch=watch_znode)
            dis = input("display? ")
            if dis == "d":
                display_tree_structure(zk)
    except KeyboardInterrupt:
        pass

    zk.stop()


if __name__ == "__main__":
    main()
