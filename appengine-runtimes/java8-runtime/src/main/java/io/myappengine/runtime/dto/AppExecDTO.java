package io.myappengine.runtime.dto;

import java.util.List;

public class AppExecDTO {

    private List<ParamDTO> queryParams;
    private List<ParamDTO> pathParams;
    private Object         payload;

    public List<ParamDTO> getQueryParams() {
        return queryParams;
    }

    public void setQueryParams(List<ParamDTO> queryParams) {
        this.queryParams = queryParams;
    }

    public List<ParamDTO> getPathParams() {
        return pathParams;
    }

    public void setPathParams(List<ParamDTO> pathParams) {
        this.pathParams = pathParams;
    }

    public Object getPayload() {
        return payload;
    }

    public void setPayload(Object payload) {
        this.payload = payload;
    }

}
