


  //TODO 代码
    %s := &virhis.%s{
        %s
    }
    %s := new(virhis.%s)
    err := his.%s(%s, %s)
    if err != nil {
        logger.Error("call %s err, please remote", err)
        moduleman.NewError(std.AdminCodeServerError, err.Error())
    }



    %s