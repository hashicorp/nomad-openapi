/*
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.nomadproject.client.api;

import io.nomadproject.client.ApiException;
import io.nomadproject.client.models.FuzzySearchRequest;
import io.nomadproject.client.models.FuzzySearchResponse;
import io.nomadproject.client.models.SearchRequest;
import io.nomadproject.client.models.SearchResponse;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for SearchApi
 */
@Ignore
public class SearchApiTest {

    private final SearchApi api = new SearchApi();

    
    /**
     * 
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getFuzzySearchTest() throws ApiException {
        FuzzySearchRequest fuzzySearchRequest = null;
        String region = null;
        String namespace = null;
        Integer index = null;
        String wait = null;
        String stale = null;
        String prefix = null;
        String xNomadToken = null;
        Integer perPage = null;
        String nextToken = null;
                FuzzySearchResponse response = api.getFuzzySearch(fuzzySearchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
        // TODO: test validations
    }
    
    /**
     * 
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSearchTest() throws ApiException {
        SearchRequest searchRequest = null;
        String region = null;
        String namespace = null;
        Integer index = null;
        String wait = null;
        String stale = null;
        String prefix = null;
        String xNomadToken = null;
        Integer perPage = null;
        String nextToken = null;
                SearchResponse response = api.getSearch(searchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
        // TODO: test validations
    }
    
}
