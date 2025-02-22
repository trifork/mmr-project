/* tslint:disable */
/* eslint-disable */
/**
 * MMRProject.Api
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  ClaimProfileRequest,
  ProfileDetails,
} from '../models/index';
import {
    ClaimProfileRequestFromJSON,
    ClaimProfileRequestToJSON,
    ProfileDetailsFromJSON,
    ProfileDetailsToJSON,
} from '../models/index';

export interface ProfileClaimProfileRequest {
    claimProfileRequest: ClaimProfileRequest;
}

/**
 * 
 */
export class ProfileApi extends runtime.BaseAPI {

    /**
     */
    async profileClaimProfileRaw(requestParameters: ProfileClaimProfileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ProfileDetails>> {
        if (requestParameters['claimProfileRequest'] == null) {
            throw new runtime.RequiredError(
                'claimProfileRequest',
                'Required parameter "claimProfileRequest" was null or undefined when calling profileClaimProfile().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v1/profile/claim`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ClaimProfileRequestToJSON(requestParameters['claimProfileRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ProfileDetailsFromJSON(jsonValue));
    }

    /**
     */
    async profileClaimProfile(requestParameters: ProfileClaimProfileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ProfileDetails> {
        const response = await this.profileClaimProfileRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async profileGetProfileRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ProfileDetails>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/profile`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ProfileDetailsFromJSON(jsonValue));
    }

    /**
     */
    async profileGetProfile(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ProfileDetails> {
        const response = await this.profileGetProfileRaw(initOverrides);
        return await response.value();
    }

}
