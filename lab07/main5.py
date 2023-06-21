import subprocess
import threading
from time import sleep
from kazoo.client import KazooClient
from kazoo.recipe.watchers import ChildrenWatch

GRAPHICAL_APP_PATH = "/usr/bin/gedit"
process = None
t1 = None

def my_job():
    print("serio startuje się apka")
    process = subprocess.Popen([GRAPHICAL_APP_PATH, "&"], shell=True)

def start_external_app():
    try:
        t1=threading.Thread(target=my_job)
        t1.start()
        sleep(1)
        print(f"created process: {process.pid}")
    except:
        print("no nie odpala się")

def stop_external_app():
    try:
        process.terminate()
    except:
        print("no nie zamyka się")


def display_num_children(children):
    num_children = len(children)
    print("Aktualna ilość potomków: {}".format(num_children))


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

    ChildrenWatch(client=zk, path="/z", func=watch_children)
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
