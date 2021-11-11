// package: heyrenee.v1
// file: heyrenee/v1/rpm_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_rpm_schedule_pb from "../../heyrenee/v1/messages/rpm_schedule_pb";
import * as heyrenee_v1_messages_rpm_measurement_pb from "../../heyrenee/v1/messages/rpm_measurement_pb";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";

export class ListRpmSchedulesRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getRpmScheduleStatus(): heyrenee_v1_messages_rpm_schedule_pb.RpmScheduleStatusMap[keyof heyrenee_v1_messages_rpm_schedule_pb.RpmScheduleStatusMap];
  setRpmScheduleStatus(value: heyrenee_v1_messages_rpm_schedule_pb.RpmScheduleStatusMap[keyof heyrenee_v1_messages_rpm_schedule_pb.RpmScheduleStatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRpmSchedulesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRpmSchedulesRequest): ListRpmSchedulesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRpmSchedulesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRpmSchedulesRequest;
  static deserializeBinaryFromReader(message: ListRpmSchedulesRequest, reader: jspb.BinaryReader): ListRpmSchedulesRequest;
}

export namespace ListRpmSchedulesRequest {
  export type AsObject = {
    patientId: string,
    rpmScheduleStatus: heyrenee_v1_messages_rpm_schedule_pb.RpmScheduleStatusMap[keyof heyrenee_v1_messages_rpm_schedule_pb.RpmScheduleStatusMap],
  }
}

export class ListRpmSchedulesResponse extends jspb.Message {
  clearRpmSchedulesList(): void;
  getRpmSchedulesList(): Array<heyrenee_v1_messages_rpm_schedule_pb.RpmSchedule>;
  setRpmSchedulesList(value: Array<heyrenee_v1_messages_rpm_schedule_pb.RpmSchedule>): void;
  addRpmSchedules(value?: heyrenee_v1_messages_rpm_schedule_pb.RpmSchedule, index?: number): heyrenee_v1_messages_rpm_schedule_pb.RpmSchedule;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRpmSchedulesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRpmSchedulesResponse): ListRpmSchedulesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRpmSchedulesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRpmSchedulesResponse;
  static deserializeBinaryFromReader(message: ListRpmSchedulesResponse, reader: jspb.BinaryReader): ListRpmSchedulesResponse;
}

export namespace ListRpmSchedulesResponse {
  export type AsObject = {
    rpmSchedulesList: Array<heyrenee_v1_messages_rpm_schedule_pb.RpmSchedule.AsObject>,
  }
}

export class ListRpmMeasurementsRequest extends jspb.Message {
  getRpmScheduleId(): string;
  setRpmScheduleId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRpmMeasurementsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRpmMeasurementsRequest): ListRpmMeasurementsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRpmMeasurementsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRpmMeasurementsRequest;
  static deserializeBinaryFromReader(message: ListRpmMeasurementsRequest, reader: jspb.BinaryReader): ListRpmMeasurementsRequest;
}

export namespace ListRpmMeasurementsRequest {
  export type AsObject = {
    rpmScheduleId: string,
  }
}

export class ListRpmMeasurementsResponse extends jspb.Message {
  clearRpmMeasurementsList(): void;
  getRpmMeasurementsList(): Array<heyrenee_v1_messages_rpm_measurement_pb.RpmMeasurement>;
  setRpmMeasurementsList(value: Array<heyrenee_v1_messages_rpm_measurement_pb.RpmMeasurement>): void;
  addRpmMeasurements(value?: heyrenee_v1_messages_rpm_measurement_pb.RpmMeasurement, index?: number): heyrenee_v1_messages_rpm_measurement_pb.RpmMeasurement;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRpmMeasurementsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRpmMeasurementsResponse): ListRpmMeasurementsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRpmMeasurementsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRpmMeasurementsResponse;
  static deserializeBinaryFromReader(message: ListRpmMeasurementsResponse, reader: jspb.BinaryReader): ListRpmMeasurementsResponse;
}

export namespace ListRpmMeasurementsResponse {
  export type AsObject = {
    rpmMeasurementsList: Array<heyrenee_v1_messages_rpm_measurement_pb.RpmMeasurement.AsObject>,
  }
}

