import signal
import sys
import time
import Ice

Ice.loadSlice("-I. --all ../slice/zadi4.ice")
import ZadI4

def run(communicator):
    print("running")
    twoway = ZadI4.TestingServicePrx.checkedCast(
        communicator.propertyToProxy('TestingService.Proxy').ice_twoway().ice_secure(False))
    if not twoway:
        print("invalid proxy")
        sys.exit(1)

with Ice.initialize(sys.argv, "../config.client") as communicator:

    if len(sys.argv) > 1:
        print(sys.argv[0] + ": too many arguments")
        sys.exit(1)

    run(communicator)
