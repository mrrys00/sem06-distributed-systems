//
// Copyright (c) ZeroC, Inc. All rights reserved.
//
//
// Ice version 3.7.9
//
// <auto-generated>
//
// Generated from file `i4ice.ice'
//
// Warning: do not edit this file.
//
// </auto-generated>
//

package OperationsWithOptionals;

public class Time extends com.zeroc.Ice.Value
{
    public Time()
    {
    }

    public Time(int hours, int minutes)
    {
        this.hours = hours;
        this.minutes = minutes;
    }

    public Time(int hours, int minutes, int seconds)
    {
        this.hours = hours;
        this.minutes = minutes;
        setSeconds(seconds);
    }

    public int hours;

    public int minutes;

    private int seconds;
    private boolean _seconds;

    public int getSeconds()
    {
        if(!_seconds)
        {
            throw new java.util.NoSuchElementException("seconds is not set");
        }
        return seconds;
    }

    public void setSeconds(int seconds)
    {
        _seconds = true;
        this.seconds = seconds;
    }

    public boolean hasSeconds()
    {
        return _seconds;
    }

    public void clearSeconds()
    {
        _seconds = false;
    }

    public void optionalSeconds(java.util.OptionalInt v)
    {
        if(v == null || !v.isPresent())
        {
            _seconds = false;
        }
        else
        {
            _seconds = true;
            seconds = v.getAsInt();
        }
    }

    public java.util.OptionalInt optionalSeconds()
    {
        if(_seconds)
        {
            return java.util.OptionalInt.of(seconds);
        }
        else
        {
            return java.util.OptionalInt.empty();
        }
    }

    public Time clone()
    {
        return (Time)super.clone();
    }

    public static String ice_staticId()
    {
        return "::OperationsWithOptionals::Time";
    }

    @Override
    public String ice_id()
    {
        return ice_staticId();
    }

    /** @hidden */
    public static final long serialVersionUID = 5645246554979165015L;

    /** @hidden */
    @Override
    protected void _iceWriteImpl(com.zeroc.Ice.OutputStream ostr_)
    {
        ostr_.startSlice(ice_staticId(), -1, true);
        ostr_.writeInt(hours);
        ostr_.writeInt(minutes);
        if(_seconds)
        {
            ostr_.writeInt(3, seconds);
        }
        ostr_.endSlice();
    }

    /** @hidden */
    @Override
    protected void _iceReadImpl(com.zeroc.Ice.InputStream istr_)
    {
        istr_.startSlice();
        hours = istr_.readInt();
        minutes = istr_.readInt();
        if(_seconds = istr_.readOptional(3, com.zeroc.Ice.OptionalFormat.F4))
        {
            seconds = istr_.readInt();
        }
        istr_.endSlice();
    }
}