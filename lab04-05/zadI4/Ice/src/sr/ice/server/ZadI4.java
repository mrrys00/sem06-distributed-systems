package sr.ice.server;
// **********************************************************************
//
// Copyright (c) 2003-2019 ZeroC, Inc. All rights reserved.
//
// This copy of Ice is licensed to you under the terms described in the
// ICE_LICENSE file included in this distribution.
//
// **********************************************************************

import ZadI4.TestingService;
import ZadI4.Person;
import com.zeroc.Ice.Current;

public class ZadI4 implements TestingService {
    private static final int Sleep = 4000;

    @Override
    public void TestingOperation(Person person, Current current) {
        try {
            Thread.sleep(Sleep);
            System.out.println("request");
            TestResults(person);
        } catch (Exception e) {
            System.out.println(e);
        }
    }

    private void TestResults(Person person) {
        String firstname = "firstname: " + person.firstName;
        String middlename = "";
        String lasename = "lasename: " + person.lastName;
        String birthdate = "";

        if (person.hasMiddleName()) {
            middlename = "middlename: " + person.getMiddleName();
        }
        if (person.hasBirthDate()) {
            birthdate = "birthdate: " + person.getBirthDate();
        }

        System.out.println(firstname);
        System.out.println(middlename);
        System.out.println(lasename);
        System.out.println(birthdate);

    }
}
