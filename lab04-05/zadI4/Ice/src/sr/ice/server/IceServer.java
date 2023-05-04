package sr.ice.server;
// **********************************************************************
//
// Copyright (c) 2003-2019 ZeroC, Inc. All rights reserved.
//
// This copy of Ice is licensed to you under the terms described in the
// ICE_LICENSE file included in this distribution.
//
// **********************************************************************

import com.zeroc.Ice.Communicator;
import com.zeroc.Ice.Identity;
import com.zeroc.Ice.Object;
import com.zeroc.Ice.ObjectAdapter;

import static com.zeroc.Ice.Util.initialize;

public class IceServer
{
	public static void t1(String[] args)
	{
		int status = 0;
		Communicator communicator = null;

		try	{
			communicator = initialize(args);
			ObjectAdapter adapter = communicator.createObjectAdapter("Adapter1");

			ZadI4 zadI4 = new ZadI4();

			adapter.add((Object) zadI4, new Identity("zadI41", "zadI4"));
			adapter.activate();

			System.out.println("Działa ale jakim kosztem...");

			communicator.waitForShutdown();
		}

		catch (Exception e) {
			System.err.println(e);
			status = 1;
		}

		if (communicator != null) {
			try {
				communicator.destroy();
			} catch (Exception e) {
				System.err.println(e);
				status = 1;
			}
		}

		System.exit(status);
	}

	public static void main(String[] args)
	{
		IceServer app = new IceServer();
		app.t1(args);
	}
}
