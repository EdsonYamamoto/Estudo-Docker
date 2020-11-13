package com.example.demo.controller;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.List;

import javax.servlet.http.HttpServletRequest;

import ch.qos.logback.core.net.server.Client;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.example.demo.model.Cliente;

@RestController
public class ClienteController {

    @RequestMapping(value = "/cliente", method = RequestMethod.GET)
    public Cliente cliente() {
        return new Cliente();
    }

    @RequestMapping(value = "/listaClientes", method = RequestMethod.GET)
    public List<Cliente> listarClientes( ) {
        List<Cliente> clientes = new ArrayList<Cliente>();
        clientes.add(new Cliente());
        clientes.add(new Cliente());
        clientes.add(new Cliente());

        return clientes;
    }

}