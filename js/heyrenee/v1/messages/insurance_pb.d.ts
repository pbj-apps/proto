// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/insurance.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";

export class Insurance extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getInsuranceId(): string;
  setInsuranceId(value: string): void;

  getInsurer(): string;
  setInsurer(value: string): void;

  getInsuranceType(): InsuranceTypeMap[keyof InsuranceTypeMap];
  setInsuranceType(value: InsuranceTypeMap[keyof InsuranceTypeMap]): void;

  getInsuranceStatus(): InsuranceStatusMap[keyof InsuranceStatusMap];
  setInsuranceStatus(value: InsuranceStatusMap[keyof InsuranceStatusMap]): void;

  getPolicyOwnerName(): string;
  setPolicyOwnerName(value: string): void;

  clearPolicyOwnerAddressLinesList(): void;
  getPolicyOwnerAddressLinesList(): Array<string>;
  setPolicyOwnerAddressLinesList(value: Array<string>): void;
  addPolicyOwnerAddressLines(value: string, index?: number): string;

  getPolicyOwnerCity(): string;
  setPolicyOwnerCity(value: string): void;

  getPolicyOwnerState(): string;
  setPolicyOwnerState(value: string): void;

  getPolicyOwnerZip(): string;
  setPolicyOwnerZip(value: string): void;

  getPolicyOwnerPhone(): string;
  setPolicyOwnerPhone(value: string): void;

  getGroupNumber(): string;
  setGroupNumber(value: string): void;

  getPolicyNumber(): string;
  setPolicyNumber(value: string): void;

  getRxPolicyNumber(): string;
  setRxPolicyNumber(value: string): void;

  clearClaimsAddressLinesList(): void;
  getClaimsAddressLinesList(): Array<string>;
  setClaimsAddressLinesList(value: Array<string>): void;
  addClaimsAddressLines(value: string, index?: number): string;

  getClaimsCity(): string;
  setClaimsCity(value: string): void;

  getClaimsState(): string;
  setClaimsState(value: string): void;

  getClaimsZip(): string;
  setClaimsZip(value: string): void;

  getClaimsPhone(): string;
  setClaimsPhone(value: string): void;

  clearRxClaimsAddressLinesList(): void;
  getRxClaimsAddressLinesList(): Array<string>;
  setRxClaimsAddressLinesList(value: Array<string>): void;
  addRxClaimsAddressLines(value: string, index?: number): string;

  getRxClaimsCity(): string;
  setRxClaimsCity(value: string): void;

  getRxClaimsState(): string;
  setRxClaimsState(value: string): void;

  getRxClaimsZip(): string;
  setRxClaimsZip(value: string): void;

  getRxClaimsPhone(): string;
  setRxClaimsPhone(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Insurance.AsObject;
  static toObject(includeInstance: boolean, msg: Insurance): Insurance.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Insurance, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Insurance;
  static deserializeBinaryFromReader(message: Insurance, reader: jspb.BinaryReader): Insurance;
}

export namespace Insurance {
  export type AsObject = {
    patientId: string,
    insuranceId: string,
    insurer: string,
    insuranceType: InsuranceTypeMap[keyof InsuranceTypeMap],
    insuranceStatus: InsuranceStatusMap[keyof InsuranceStatusMap],
    policyOwnerName: string,
    policyOwnerAddressLinesList: Array<string>,
    policyOwnerCity: string,
    policyOwnerState: string,
    policyOwnerZip: string,
    policyOwnerPhone: string,
    groupNumber: string,
    policyNumber: string,
    rxPolicyNumber: string,
    claimsAddressLinesList: Array<string>,
    claimsCity: string,
    claimsState: string,
    claimsZip: string,
    claimsPhone: string,
    rxClaimsAddressLinesList: Array<string>,
    rxClaimsCity: string,
    rxClaimsState: string,
    rxClaimsZip: string,
    rxClaimsPhone: string,
  }
}

export interface InsuranceTypeMap {
  INSURANCE_TYPE_UNSPECIFIED: 0;
  INSURANCE_TYPE_MEDICAID: 1;
  INSURANCE_TYPE_MEDICARE: 2;
  INSURANCE_TYPE_MEDICARE_ADVANTAGE: 3;
  INSURANCE_TYPE_VETERANS_AFFAIRS: 4;
  INSURANCE_TYPE_EMPLOYER_BASED: 5;
  INSURANCE_TYPE_PRIVATE: 6;
}

export const InsuranceType: InsuranceTypeMap;

export interface InsuranceStatusMap {
  INSURANCE_STATUS_UNSPECIFIED: 0;
  INSURANCE_STATUS_ACTIVE: 1;
  INSURANCE_STATUS_INACTIVE: 2;
}

export const InsuranceStatus: InsuranceStatusMap;

