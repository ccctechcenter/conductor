/**
 * Copyright 2016 Netflix, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/**
 * 
 */
package com.netflix.conductor.contribs.http;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;
import com.netflix.conductor.common.metadata.tasks.Task;
import com.netflix.conductor.common.metadata.tasks.Task.Status;
import com.netflix.conductor.common.metadata.workflow.TaskType;
import com.netflix.conductor.common.metadata.workflow.WorkflowDef;
import com.netflix.conductor.common.metadata.workflow.WorkflowTask;
import com.netflix.conductor.common.run.Workflow;
import com.netflix.conductor.contribs.http.CcctcHttpTask;
import com.netflix.conductor.contribs.http.HttpTask.Input;
import com.netflix.conductor.core.config.Configuration;
import com.netflix.conductor.core.execution.DeciderService;
import com.netflix.conductor.core.execution.ParametersUtils;
import com.netflix.conductor.core.execution.WorkflowExecutor;
import com.netflix.conductor.core.execution.mapper.TaskMapper;
import com.netflix.conductor.core.utils.ExternalPayloadStorageUtils;
import com.netflix.conductor.dao.MetadataDAO;
import com.netflix.conductor.dao.QueueDAO;
import java.time.Instant;
import org.eclipse.jetty.server.Request;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.handler.AbstractHandler;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.junit.*;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;
import org.mockito.Mockito;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertTrue;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

/**
 * @author jstanley
 *
 */
@SuppressWarnings("unchecked")
public class TestCcctcHttpTask {

    private static final String ERROR_RESPONSE = "Something went wrong!";
    
    private static final String TEXT_RESPONSE = "Text Response";
    
    private static final double NUM_RESPONSE = 42.42d;
    
    private static String JSON_RESPONSE;
    
    private CcctcHttpTask httpTask;
    
    private WorkflowExecutor workflowExecutor;
    private Configuration config;

    private Workflow workflow = new Workflow();
    
    private static Server server;
    
    private static ObjectMapper objectMapper = new ObjectMapper();
   
    @BeforeClass
    public static void init() throws Exception {

        Map<String, Object> map = new HashMap<>();
        map.put("key", "value1");
        map.put("num", 42);
        JSON_RESPONSE = objectMapper.writeValueAsString(map);

        server = new Server(7010);
        ServletContextHandler servletContextHandler = new ServletContextHandler(server, "/", ServletContextHandler.SESSIONS);
        servletContextHandler.setHandler(new EchoHandler());
        server.start();
    }

    @AfterClass
    public static void cleanup() {
        if (server != null) {
            try {
                server.stop();
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }

    @Before
    public void setup() {
        workflowExecutor = mock(WorkflowExecutor.class);
        config = mock(Configuration.class);
        RestClientManager rcm = new RestClientManager(Mockito.mock(Configuration.class));
        when(config.getServerId()).thenReturn("test_server_id");
        httpTask = new CcctcHttpTask(rcm, config);
    }

    @Test
    public void testRequestHandled() throws Exception {
        Task task = new Task();
        Input input = new Input();
        task.getInputData().put(HttpTask.REQUEST_PARAMETER_NAME, input);
        JsonNode status = objectMapper.readTree("{\"400\":\"COMPLETED\"}");
        task.getInputData().put(CcctcHttpTask.HTTP_STATUS_OVERIDE_PARAMETER_NAME, status);
        HttpTask.HttpResponse response = new HttpTask.HttpResponse();
        response.statusCode = 400;
        assertTrue("Should response handled be true", httpTask.handleOptionalResponse(task, response));
        assertEquals("Should task status set to COMPLETED", task.getStatus(),Status.COMPLETED);
        assertEquals("Should task output contans httpStatus 400", task.getOutputData().get("httpStatus"),"400");
        assertTrue("Should response handled be true",(Boolean)task.getOutputData().get("overrideActivated"));
    }

    @Test
    public void testRequestNotHandled() throws Exception {
        Task task = new Task();
        Input input = new Input();
        task.getInputData().put(HttpTask.REQUEST_PARAMETER_NAME, input);
        JsonNode status = objectMapper.readTree("{\"400\":\"COMPLETED\"}");
        task.getInputData().put(CcctcHttpTask.HTTP_STATUS_OVERIDE_PARAMETER_NAME, status);
        HttpTask.HttpResponse response = new HttpTask.HttpResponse();
        response.statusCode = 403;
        assertFalse("Should response handled be false", httpTask.handleOptionalResponse(task, response));
        assertNull("Should task status set to null", task.getStatus());
        assertFalse("Should response handled be false",(Boolean)task.getOutputData().get("overrideActivated"));
    }

    private static class EchoHandler extends AbstractHandler {

        private TypeReference<Map<String, Object>> mapOfObj = new TypeReference<Map<String, Object>>() {
        };

        @Override
        public void handle(String target, Request baseRequest, HttpServletRequest request, HttpServletResponse response)
                throws IOException {
            if (request.getMethod().equals("GET") && request.getRequestURI().equals("/text")) {
                PrintWriter writer = response.getWriter();
                writer.print(TEXT_RESPONSE);
                writer.flush();
                writer.close();
            } else if (request.getMethod().equals("GET") && request.getRequestURI().equals("/json")) {
                response.addHeader("Content-Type", "application/json");
                PrintWriter writer = response.getWriter();
                writer.print(JSON_RESPONSE);
                writer.flush();
                writer.close();
            } else if (request.getMethod().equals("GET") && request.getRequestURI().equals("/failure")) {
                response.addHeader("Content-Type", "text/plain");
                response.setStatus(500);
                PrintWriter writer = response.getWriter();
                writer.print(ERROR_RESPONSE);
                writer.flush();
                writer.close();
            } else if (request.getMethod().equals("POST") && request.getRequestURI().equals("/post")) {
                response.addHeader("Content-Type", "application/json");
                BufferedReader reader = request.getReader();
                Map<String, Object> input = objectMapper.readValue(reader, mapOfObj);
                Set<String> keys = input.keySet();
                for (String key : keys) {
                    input.put(key, key);
                }
                PrintWriter writer = response.getWriter();
                writer.print(objectMapper.writeValueAsString(input));
                writer.flush();
                writer.close();
            } else if (request.getMethod().equals("POST") && request.getRequestURI().equals("/post2")) {
                response.addHeader("Content-Type", "application/json");
                response.setStatus(204);
                BufferedReader reader = request.getReader();
                Map<String, Object> input = objectMapper.readValue(reader, mapOfObj);
                Set<String> keys = input.keySet();
                System.out.println(keys);
                response.getWriter().close();

            } else if (request.getMethod().equals("GET") && request.getRequestURI().equals("/numeric")) {
                PrintWriter writer = response.getWriter();
                writer.print(NUM_RESPONSE);
                writer.flush();
                writer.close();
            } else if (request.getMethod().equals("POST") && request.getRequestURI().equals("/oauth")) {
                //echo back oauth parameters generated in the Authorization header in the response
                Map<String, String> params = parseOauthParameters(request);
                response.addHeader("Content-Type", "application/json");
                PrintWriter writer = response.getWriter();
                writer.print(objectMapper.writeValueAsString(params));
                writer.flush();
                writer.close();
            }
        }

        private Map<String, String> parseOauthParameters(HttpServletRequest request) {
            String paramString = request.getHeader("Authorization").replaceAll("^OAuth (.*)", "$1");
            return Arrays.stream(paramString.split("\\s*,\\s*"))
                    .map(pair -> pair.split("="))
                    .collect(Collectors.toMap(o -> o[0], o -> o[1].replaceAll("\"", "")));
        }
    }
}
    