
#ifndef CALC_ICE
#define CALC_ICE

module ZadI4 {
    enum EEnum {
        CYAN,
        YELLOW,
        MAGENTA
    };

    exception NoInput {};

    class Time {
        int hours;
        int minutes;
        optional(3) int seconds;
    };

    class Request {
        int intArg;
        string strArg;
        EEnum enumArg;
        Time timeArg;

        optional(5) int optIntArg;
        optional(6) string optStrArg;
        optional(7) EEnum optEnumArg;
        optional(8) Time optTimeArg;
    };

    class Response {
        string strResp;
        EEnum enumArg;

        optional(3) string optStrResp;
        optional(4) EEnum optEnumArg;
    };

    interface ExampleService {
        Response OppOperation(Request request);
    };
};
#endif
