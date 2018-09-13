

    //TODO 代码
    %s := &virhis.%s{
        //TODO 请求参数
    }
    %s := new(virhis.%s)
    err = his.%s(%s, %s)
    if err != nil {
        logger.Error("call %s err, please remote", err)
        moduleman.NewError(std.AdminCodeServerError, err.Error())
    }

    list := make([]%s, 0, len(%s.List))
    for _, v := range %s.List {
        list = append(list, %s{
            //TODO 属性
        })
    }

    reply.List = list