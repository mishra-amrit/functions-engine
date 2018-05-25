/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
