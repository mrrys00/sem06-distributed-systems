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

package ZadI4;

public class Request extends com.zeroc.Ice.Value
{
    public Request()
    {
        this.strArg = "";
        this.enumArg = EEnum.CYAN;
        this.optStrArg = "";
        this.optEnumArg = EEnum.CYAN;
    }

    public Request(int intArg, String strArg, EEnum enumArg, Time timeArg)
    {
        this.intArg = intArg;
        this.strArg = strArg;
        this.enumArg = enumArg;
        this.timeArg = timeArg;
        this.optStrArg = "";
        this.optEnumArg = EEnum.CYAN;
    }

    public Request(int intArg, String strArg, EEnum enumArg, Time timeArg, int optIntArg, String optStrArg, EEnum optEnumArg, Time optTimeArg)
    {
        this.intArg = intArg;
        this.strArg = strArg;
        this.enumArg = enumArg;
        this.timeArg = timeArg;
        setOptIntArg(optIntArg);
        setOptStrArg(optStrArg);
        setOptEnumArg(optEnumArg);
        setOptTimeArg(optTimeArg);
    }

    public int intArg;

    public String strArg;

    public EEnum enumArg;

    public Time timeArg;

    private int optIntArg;
    private boolean _optIntArg;

    public int getOptIntArg()
    {
        if(!_optIntArg)
        {
            throw new java.util.NoSuchElementException("optIntArg is not set");
        }
        return optIntArg;
    }

    public void setOptIntArg(int optIntArg)
    {
        _optIntArg = true;
        this.optIntArg = optIntArg;
    }

    public boolean hasOptIntArg()
    {
        return _optIntArg;
    }

    public void clearOptIntArg()
    {
        _optIntArg = false;
    }

    public void optionalOptIntArg(java.util.OptionalInt v)
    {
        if(v == null || !v.isPresent())
        {
            _optIntArg = false;
        }
        else
        {
            _optIntArg = true;
            optIntArg = v.getAsInt();
        }
    }

    public java.util.OptionalInt optionalOptIntArg()
    {
        if(_optIntArg)
        {
            return java.util.OptionalInt.of(optIntArg);
        }
        else
        {
            return java.util.OptionalInt.empty();
        }
    }

    private String optStrArg;
    private boolean _optStrArg;

    public String getOptStrArg()
    {
        if(!_optStrArg)
        {
            throw new java.util.NoSuchElementException("optStrArg is not set");
        }
        return optStrArg;
    }

    public void setOptStrArg(String optStrArg)
    {
        _optStrArg = true;
        this.optStrArg = optStrArg;
    }

    public boolean hasOptStrArg()
    {
        return _optStrArg;
    }

    public void clearOptStrArg()
    {
        _optStrArg = false;
    }

    public void optionalOptStrArg(java.util.Optional<java.lang.String> v)
    {
        if(v == null || !v.isPresent())
        {
            _optStrArg = false;
        }
        else
        {
            _optStrArg = true;
            optStrArg = v.get();
        }
    }

    public java.util.Optional<java.lang.String> optionalOptStrArg()
    {
        if(_optStrArg)
        {
            return java.util.Optional.of(optStrArg);
        }
        else
        {
            return java.util.Optional.empty();
        }
    }

    private EEnum optEnumArg;
    private boolean _optEnumArg;

    public EEnum getOptEnumArg()
    {
        if(!_optEnumArg)
        {
            throw new java.util.NoSuchElementException("optEnumArg is not set");
        }
        return optEnumArg;
    }

    public void setOptEnumArg(EEnum optEnumArg)
    {
        _optEnumArg = true;
        this.optEnumArg = optEnumArg;
    }

    public boolean hasOptEnumArg()
    {
        return _optEnumArg;
    }

    public void clearOptEnumArg()
    {
        _optEnumArg = false;
    }

    public void optionalOptEnumArg(java.util.Optional<EEnum> v)
    {
        if(v == null || !v.isPresent())
        {
            _optEnumArg = false;
        }
        else
        {
            _optEnumArg = true;
            optEnumArg = v.get();
        }
    }

    public java.util.Optional<EEnum> optionalOptEnumArg()
    {
        if(_optEnumArg)
        {
            return java.util.Optional.of(optEnumArg);
        }
        else
        {
            return java.util.Optional.empty();
        }
    }

    private Time optTimeArg;
    private boolean _optTimeArg;

    public Time getOptTimeArg()
    {
        if(!_optTimeArg)
        {
            throw new java.util.NoSuchElementException("optTimeArg is not set");
        }
        return optTimeArg;
    }

    public void setOptTimeArg(Time optTimeArg)
    {
        _optTimeArg = true;
        this.optTimeArg = optTimeArg;
    }

    public boolean hasOptTimeArg()
    {
        return _optTimeArg;
    }

    public void clearOptTimeArg()
    {
        _optTimeArg = false;
    }

    public void optionalOptTimeArg(java.util.Optional<Time> v)
    {
        if(v == null || !v.isPresent())
        {
            _optTimeArg = false;
        }
        else
        {
            _optTimeArg = true;
            optTimeArg = v.get();
        }
    }

    public java.util.Optional<Time> optionalOptTimeArg()
    {
        if(_optTimeArg)
        {
            return java.util.Optional.ofNullable(optTimeArg);
        }
        else
        {
            return java.util.Optional.empty();
        }
    }

    public Request clone()
    {
        return (Request)super.clone();
    }

    public static String ice_staticId()
    {
        return "::ZadI4::Request";
    }

    @Override
    public String ice_id()
    {
        return ice_staticId();
    }

    /** @hidden */
    public static final long serialVersionUID = -1821715485574006398L;

    /** @hidden */
    @Override
    protected void _iceWriteImpl(com.zeroc.Ice.OutputStream ostr_)
    {
        ostr_.startSlice(ice_staticId(), -1, true);
        ostr_.writeInt(intArg);
        ostr_.writeString(strArg);
        EEnum.ice_write(ostr_, enumArg);
        ostr_.writeValue(timeArg);
        if(_optIntArg)
        {
            ostr_.writeInt(5, optIntArg);
        }
        if(_optStrArg)
        {
            ostr_.writeString(6, optStrArg);
        }
        if(_optEnumArg)
        {
            EEnum.ice_write(ostr_, 7, optEnumArg);
        }
        if(_optTimeArg)
        {
            ostr_.writeValue(8, optTimeArg);
        }
        ostr_.endSlice();
    }

    /** @hidden */
    @Override
    protected void _iceReadImpl(com.zeroc.Ice.InputStream istr_)
    {
        istr_.startSlice();
        intArg = istr_.readInt();
        strArg = istr_.readString();
        enumArg = EEnum.ice_read(istr_);
        istr_.readValue(v -> timeArg = v, Time.class);
        if(_optIntArg = istr_.readOptional(5, com.zeroc.Ice.OptionalFormat.F4))
        {
            optIntArg = istr_.readInt();
        }
        if(_optStrArg = istr_.readOptional(6, com.zeroc.Ice.OptionalFormat.VSize))
        {
            optStrArg = istr_.readString();
        }
        if(_optEnumArg = istr_.readOptional(7, com.zeroc.Ice.OptionalFormat.Size))
        {
            optEnumArg = EEnum.ice_read(istr_);
        }
        if(_optTimeArg = istr_.readOptional(8, com.zeroc.Ice.OptionalFormat.Class))
        {
            istr_.readValue(v -> optTimeArg = v, Time.class);
        }
        istr_.endSlice();
    }
}
