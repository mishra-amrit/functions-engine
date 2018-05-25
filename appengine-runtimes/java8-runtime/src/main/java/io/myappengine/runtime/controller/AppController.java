package io.myappengine.runtime.controller;

import java.lang.reflect.InvocationTargetException;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.fasterxml.jackson.databind.ObjectMapper;

import io.myappengine.runtime.core.AppManager;
import io.myappengine.runtime.core.AppRegistry;
import io.myappengine.runtime.dto.AppDTO;
import io.myappengine.runtime.dto.AppExecDTO;
import io.myappengine.runtime.dto.ResponseDTO;

@RestController
public class AppController {

    @RequestMapping(value = "/load", method = RequestMethod.POST)
    public ResponseEntity<ResponseDTO> loadApp(@RequestBody AppDTO appDTO) {

	HttpStatus responseStatus = null;
	try {
	    AppRegistry.getInstance().register(appDTO, false);
	    responseStatus = HttpStatus.OK;
	} catch (Exception e) {
	    e.printStackTrace();
	    responseStatus = HttpStatus.EXPECTATION_FAILED;
	}

	return new ResponseEntity<ResponseDTO>(new ResponseDTO(), responseStatus);
    }

    @RequestMapping(value = "/load", method = RequestMethod.PUT)
    public ResponseEntity<ResponseDTO> reloadApp(@RequestBody AppDTO appDTO) {

	HttpStatus responseStatus = null;
	try {
	    AppRegistry.getInstance().register(appDTO, true);
	    responseStatus = HttpStatus.OK;
	} catch (Exception e) {
	    e.printStackTrace();
	    responseStatus = HttpStatus.EXPECTATION_FAILED;
	}

	return new ResponseEntity<ResponseDTO>(new ResponseDTO(), responseStatus);
    }

    @RequestMapping(value = "/exec/{appName}", method = RequestMethod.POST)
    public ResponseEntity<ResponseDTO> execApp(@PathVariable("appName") String appName,
            @RequestBody AppExecDTO appExecDTO) {

	ResponseDTO responseDTO = new ResponseDTO();
	HttpStatus responseStatus = null;

	AppManager appMgr = AppRegistry.getInstance().getAppManager(appName);
	try {
	    Object execOutput = appMgr.exec(appName, appExecDTO.getPathParams(), appExecDTO.getQueryParams(),
	            appExecDTO.getPayload());
	    responseDTO.setData(execOutput);
	    responseStatus = HttpStatus.OK;
	} catch (IllegalAccessException | IllegalArgumentException | InvocationTargetException e) {
	    e.printStackTrace();
	    responseStatus = HttpStatus.EXPECTATION_FAILED;
	}

	return new ResponseEntity<ResponseDTO>(responseDTO, responseStatus);
    }

}
