package sr.ice.server;

import OperationsWithOptionals.ExampleService;
import OperationsWithOptionals.Request;
import OperationsWithOptionals.Response;
import com.zeroc.Ice.Current;

import java.util.ArrayList;
import java.util.List;

public class OperationsWithOptionals implements ExampleService {
    private static final int Sleep = 4000;

    @Override
    public Response OppOperation(Request request, Current current) {
        try {
            Thread.sleep(Sleep);
            System.out.println("client request incoming!");
            return BuildResponse(request);
        } catch (Exception e) {
            //noinspection ThrowablePrintedToSystemOut
            System.out.println(e);
            return null;
        }
    }

    private static Response BuildResponse(Request request) {
        List<String> sb1 = new ArrayList<>(), sb2 = new ArrayList<>();

        sb1.add("str: " + request.strArg);
        sb1.add("int: " + request.intArg);
        sb1.add("hh: " + request.timeArg.hours);
        sb1.add("mm: " + request.timeArg.minutes);

        if (request.timeArg.hasSeconds())
            sb1.add("ss: " + request.timeArg.getSeconds());

        if (request.hasOptStrArg() && !request.getOptStrArg().isBlank())
            sb2.add("str: " + request.getOptStrArg());

        if (request.hasOptIntArg())
            sb2.add("int: " + request.getOptIntArg());

        if (request.hasOptTimeArg()) {
            sb2.add("hh:" + request.getOptTimeArg().hours);
            sb2.add("mm:" + request.getOptTimeArg().minutes);

            if (request.getOptTimeArg().hasSeconds()) {
                sb2.add("ss:" + request.getOptTimeArg().getSeconds());
            }
        }

        var result = new Response();
        result.strResp = (String.join(", ", sb1));
        result.enumArg = request.enumArg;

        if (!sb2.isEmpty())
            result.setOptStrResp(String.join(", ", sb2));

        if (request.hasOptEnumArg())
            result.setOptEnumArg(request.getOptEnumArg());

        return result;
    }
}
