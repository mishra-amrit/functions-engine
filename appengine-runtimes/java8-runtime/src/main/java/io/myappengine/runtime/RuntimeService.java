package io.myappengine.runtime;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
public class RuntimeService {

    public static void main(String[] args) {
        SpringApplication.run(RuntimeService.class, args);
    }

}
