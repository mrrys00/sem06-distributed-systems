//
// Copyright (c) ZeroC, Inc. All rights reserved.
//
//
// Ice version 3.7.9
//
// <auto-generated>
//
// Generated from file `zadi4.ice'
//
// Warning: do not edit this file.
//
// </auto-generated>
//

package Demo;

public class NoInput extends com.zeroc.Ice.UserException
{
    public NoInput()
    {
    }

    public NoInput(Throwable cause)
    {
        super(cause);
    }

    public String ice_id()
    {
        return "::Demo::NoInput";
    }

    /** @hidden */
    @Override
    protected void _writeImpl(com.zeroc.Ice.OutputStream ostr_)
    {
        ostr_.startSlice("::Demo::NoInput", -1, true);
        ostr_.endSlice();
    }

    /** @hidden */
    @Override
    protected void _readImpl(com.zeroc.Ice.InputStream istr_)
    {
        istr_.startSlice();
        istr_.endSlice();
    }

    /** @hidden */
    public static final long serialVersionUID = -8419862277737600333L;
}
