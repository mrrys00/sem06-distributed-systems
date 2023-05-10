import signal
import sys
import time
import Ice

Ice.loadSlice("-I. --all ../slice/zadi4.ice")
import ZadI4

# from zadi4_ice import Person

# class Person(Ice.Value):
#     def __init__(self, firstName='', middleName=Ice.Unset, lastName='', birthDate=Ice.Unset):
#         self.firstName = firstName
#         self.middleName = middleName
#         self.lastName = lastName
#         self.birthDate = birthDate

def run(communicator):
    print("running")
    twoway = ZadI4.TestingServicePrx.checkedCast(
        communicator.propertyToProxy('TestingService.Proxy').ice_twoway().ice_secure(False))
    if not twoway:
        print("invalid proxy")
        sys.exit(1)

    oneway = ZadI4.TestingServicePrx.uncheckedCast(twoway.ice_oneway())
    batchOneway = ZadI4.TestingServicePrx.uncheckedCast(twoway.ice_batchOneway())
    datagram = ZadI4.TestingServicePrx.uncheckedCast(twoway.ice_datagram())
    batchDatagram = ZadI4.TestingServicePrx.uncheckedCast(twoway.ice_batchDatagram())

    secure = False
    timeout = -1
    delay = 0

    c = None
    while c != 'x':
        try:
            sys.stdout.write("==> ")
            sys.stdout.flush()
            c = sys.stdin.readline().strip()
            if c == "p1":
                p1 = ZadI4.Person("ala", "von", "neumann", 12345678)
                twoway.TestingOperation(p1)
            elif c == "p2":
                p2 = ZadI4.Person(firstName="ala", lastName="neumann", birthDate=12345678)
                oneway.TestingOperation(p2)
            elif c == "p3":
                p3 = ZadI4.Person(firstName="ala", lastName="neumann")
                oneway.TestingOperation(p3)
            elif c == "p4":
                try:
                    p4 = ZadI4.Person(lastName="neumann")
                    oneway.TestingOperation(p4)
                except:
                    print("error while creating person!")
            elif c == "x":
                break

        except Ice.Exception as ex:
            print(ex)


    print("still OK")

with Ice.initialize(sys.argv, "../config.client") as communicator:

    if len(sys.argv) > 1:
        print(sys.argv[0] + ": too many arguments")
        sys.exit(1)

    run(communicator)
