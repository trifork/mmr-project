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

import type {
  LeaderboardEntry,
  PlayerHistoryDetails,
  TimeStatisticsEntry,
} from '../models/index';
import {
  LeaderboardEntryFromJSON,
  LeaderboardEntryToJSON,
  PlayerHistoryDetailsFromJSON,
  PlayerHistoryDetailsToJSON,
  TimeStatisticsEntryFromJSON,
  TimeStatisticsEntryToJSON,
} from '../models/index';
import * as runtime from '../runtime';

export interface StatisticsGetPlayerHistoryRequest {
  userId?: number;
}

/**
 *
 */
export class StatisticsApi extends runtime.BaseAPI {
  /**
   */
  async statisticsGetLeaderboardRaw(
    initOverrides?: RequestInit | runtime.InitOverrideFunction
  ): Promise<runtime.ApiResponse<Array<LeaderboardEntry>>> {
    const queryParameters: any = {};

    const headerParameters: runtime.HTTPHeaders = {};

    const response = await this.request(
      {
        path: `/api/v1/stats/leaderboard`,
        method: 'GET',
        headers: headerParameters,
        query: queryParameters,
      },
      initOverrides
    );

    return new runtime.JSONApiResponse(response, (jsonValue) =>
      jsonValue.map(LeaderboardEntryFromJSON)
    );
  }

  /**
   */
  async statisticsGetLeaderboard(
    initOverrides?: RequestInit | runtime.InitOverrideFunction
  ): Promise<Array<LeaderboardEntry>> {
    const response = await this.statisticsGetLeaderboardRaw(initOverrides);
    return await response.value();
  }

  /**
   */
  async statisticsGetPlayerHistoryRaw(
    requestParameters: StatisticsGetPlayerHistoryRequest,
    initOverrides?: RequestInit | runtime.InitOverrideFunction
  ): Promise<runtime.ApiResponse<Array<PlayerHistoryDetails>>> {
    const queryParameters: any = {};

    if (requestParameters['userId'] != null) {
      queryParameters['userId'] = requestParameters['userId'];
    }

    const headerParameters: runtime.HTTPHeaders = {};

    const response = await this.request(
      {
        path: `/api/v1/stats/player-history`,
        method: 'GET',
        headers: headerParameters,
        query: queryParameters,
      },
      initOverrides
    );

    return new runtime.JSONApiResponse(response, (jsonValue) =>
      jsonValue.map(PlayerHistoryDetailsFromJSON)
    );
  }

  /**
   */
  async statisticsGetPlayerHistory(
    requestParameters: StatisticsGetPlayerHistoryRequest = {},
    initOverrides?: RequestInit | runtime.InitOverrideFunction
  ): Promise<Array<PlayerHistoryDetails>> {
    const response = await this.statisticsGetPlayerHistoryRaw(
      requestParameters,
      initOverrides
    );
    return await response.value();
  }

  /**
   */
  async statisticsGetTimeDistributionRaw(
    initOverrides?: RequestInit | runtime.InitOverrideFunction
  ): Promise<runtime.ApiResponse<Array<TimeStatisticsEntry>>> {
    const queryParameters: any = {};

    const headerParameters: runtime.HTTPHeaders = {};

    const response = await this.request(
      {
        path: `/api/v1/stats/time-distribution`,
        method: 'GET',
        headers: headerParameters,
        query: queryParameters,
      },
      initOverrides
    );

    return new runtime.JSONApiResponse(response, (jsonValue) =>
      jsonValue.map(TimeStatisticsEntryFromJSON)
    );
  }

  /**
   */
  async statisticsGetTimeDistribution(
    initOverrides?: RequestInit | runtime.InitOverrideFunction
  ): Promise<Array<TimeStatisticsEntry>> {
    const response = await this.statisticsGetTimeDistributionRaw(initOverrides);
    return await response.value();
  }
}
