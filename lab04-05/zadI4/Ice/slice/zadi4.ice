
#ifndef CALC_ICE
#define CALC_ICE

module ZadI4 {
    class Person {
        string firstName;
        optional(2) string middleName;
        string lastName;
        optional(1) int birthDate;
    };

    exception NoInput {};

    interface TestingService {
        void TestingOperation(Person person);
    };
};
#endif
