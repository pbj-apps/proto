// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/rpm_measurement.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class RpmMeasurement extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getRpmScheduleId(): string;
  setRpmScheduleId(value: string): void;

  getRpmMeasurementId(): string;
  setRpmMeasurementId(value: string): void;

  clearRpmMeasurementResultsList(): void;
  getRpmMeasurementResultsList(): Array<RpmMeasurementResult>;
  setRpmMeasurementResultsList(value: Array<RpmMeasurementResult>): void;
  addRpmMeasurementResults(value?: RpmMeasurementResult, index?: number): RpmMeasurementResult;

  hasTimeTaken(): boolean;
  clearTimeTaken(): void;
  getTimeTaken(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimeTaken(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RpmMeasurement.AsObject;
  static toObject(includeInstance: boolean, msg: RpmMeasurement): RpmMeasurement.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RpmMeasurement, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RpmMeasurement;
  static deserializeBinaryFromReader(message: RpmMeasurement, reader: jspb.BinaryReader): RpmMeasurement;
}

export namespace RpmMeasurement {
  export type AsObject = {
    patientId: string,
    rpmScheduleId: string,
    rpmMeasurementId: string,
    rpmMeasurementResultsList: Array<RpmMeasurementResult.AsObject>,
    timeTaken?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class RpmMeasurementResult extends jspb.Message {
  getRpmMeasurementResultId(): string;
  setRpmMeasurementResultId(value: string): void;

  getValue(): number;
  setValue(value: number): void;

  getRpmMeasurementResultUnit(): RpmMeasurementResultUnitMap[keyof RpmMeasurementResultUnitMap];
  setRpmMeasurementResultUnit(value: RpmMeasurementResultUnitMap[keyof RpmMeasurementResultUnitMap]): void;

  getRpmMeasurementResultType(): RpmMeasurementResultTypeMap[keyof RpmMeasurementResultTypeMap];
  setRpmMeasurementResultType(value: RpmMeasurementResultTypeMap[keyof RpmMeasurementResultTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RpmMeasurementResult.AsObject;
  static toObject(includeInstance: boolean, msg: RpmMeasurementResult): RpmMeasurementResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RpmMeasurementResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RpmMeasurementResult;
  static deserializeBinaryFromReader(message: RpmMeasurementResult, reader: jspb.BinaryReader): RpmMeasurementResult;
}

export namespace RpmMeasurementResult {
  export type AsObject = {
    rpmMeasurementResultId: string,
    value: number,
    rpmMeasurementResultUnit: RpmMeasurementResultUnitMap[keyof RpmMeasurementResultUnitMap],
    rpmMeasurementResultType: RpmMeasurementResultTypeMap[keyof RpmMeasurementResultTypeMap],
  }
}

export interface RpmMeasurementResultUnitMap {
  RPM_MEASUREMENT_RESULT_UNIT_UNSPECIFIED: 0;
  RPM_MEASUREMENT_RESULT_UNIT_BPM: 1;
  RPM_MEASUREMENT_RESULT_UNIT_LBS: 2;
  RPM_MEASUREMENT_RESULT_UNIT_PERCENTAGE: 3;
  RPM_MEASUREMENT_RESULT_UNIT_MG_DL: 4;
  RPM_MEASUREMENT_RESULT_UNIT_MM_HG: 5;
  RPM_MEASUREMENT_RESULT_UNIT_CELSIUS: 6;
}

export const RpmMeasurementResultUnit: RpmMeasurementResultUnitMap;

export interface RpmMeasurementResultTypeMap {
  RPM_MEASUREMENT_RESULT_TYPE_UNSPECIFIED: 0;
  RPM_MEASUREMENT_RESULT_TYPE_WEIGHT: 1;
  RPM_MEASUREMENT_RESULT_TYPE_SPO2: 2;
  RPM_MEASUREMENT_RESULT_TYPE_BPM: 3;
  RPM_MEASUREMENT_RESULT_TYPE_GLUCOSE: 4;
  RPM_MEASUREMENT_RESULT_TYPE_DIA: 5;
  RPM_MEASUREMENT_RESULT_TYPE_SYS: 6;
  RPM_MEASUREMENT_RESULT_TYPE_TEMP: 7;
}

export const RpmMeasurementResultType: RpmMeasurementResultTypeMap;

