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

import { mapValues } from '../runtime';
import type { ActiveMatchTeamDto } from './ActiveMatchTeamDto';
import {
    ActiveMatchTeamDtoFromJSON,
    ActiveMatchTeamDtoFromJSONTyped,
    ActiveMatchTeamDtoToJSON,
} from './ActiveMatchTeamDto';

/**
 * 
 * @export
 * @interface ActiveMatchDto
 */
export interface ActiveMatchDto {
    /**
     * 
     * @type {string}
     * @memberof ActiveMatchDto
     */
    id: string;
    /**
     * 
     * @type {Date}
     * @memberof ActiveMatchDto
     */
    createdAt: Date;
    /**
     * 
     * @type {ActiveMatchTeamDto}
     * @memberof ActiveMatchDto
     */
    team1: ActiveMatchTeamDto;
    /**
     * 
     * @type {ActiveMatchTeamDto}
     * @memberof ActiveMatchDto
     */
    team2: ActiveMatchTeamDto;
}

/**
 * Check if a given object implements the ActiveMatchDto interface.
 */
export function instanceOfActiveMatchDto(value: object): boolean {
    if (!('id' in value)) return false;
    if (!('createdAt' in value)) return false;
    if (!('team1' in value)) return false;
    if (!('team2' in value)) return false;
    return true;
}

export function ActiveMatchDtoFromJSON(json: any): ActiveMatchDto {
    return ActiveMatchDtoFromJSONTyped(json, false);
}

export function ActiveMatchDtoFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActiveMatchDto {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'],
        'createdAt': (new Date(json['createdAt'])),
        'team1': ActiveMatchTeamDtoFromJSON(json['team1']),
        'team2': ActiveMatchTeamDtoFromJSON(json['team2']),
    };
}

export function ActiveMatchDtoToJSON(value?: ActiveMatchDto | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'createdAt': ((value['createdAt']).toISOString()),
        'team1': ActiveMatchTeamDtoToJSON(value['team1']),
        'team2': ActiveMatchTeamDtoToJSON(value['team2']),
    };
}

