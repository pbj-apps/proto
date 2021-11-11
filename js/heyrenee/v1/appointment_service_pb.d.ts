// package: heyrenee.v1
// file: heyrenee/v1/appointment_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_appointment_pb from "../../heyrenee/v1/messages/appointment_pb";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";

export class ListAppointmentsRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getAppointmentType(): AppointmentTypeMap[keyof AppointmentTypeMap];
  setAppointmentType(value: AppointmentTypeMap[keyof AppointmentTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAppointmentsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListAppointmentsRequest): ListAppointmentsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListAppointmentsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAppointmentsRequest;
  static deserializeBinaryFromReader(message: ListAppointmentsRequest, reader: jspb.BinaryReader): ListAppointmentsRequest;
}

export namespace ListAppointmentsRequest {
  export type AsObject = {
    patientId: string,
    appointmentType: AppointmentTypeMap[keyof AppointmentTypeMap],
  }
}

export class ListAppointmentsResponse extends jspb.Message {
  clearAppointmentsList(): void;
  getAppointmentsList(): Array<heyrenee_v1_messages_appointment_pb.Appointment>;
  setAppointmentsList(value: Array<heyrenee_v1_messages_appointment_pb.Appointment>): void;
  addAppointments(value?: heyrenee_v1_messages_appointment_pb.Appointment, index?: number): heyrenee_v1_messages_appointment_pb.Appointment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAppointmentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListAppointmentsResponse): ListAppointmentsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListAppointmentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAppointmentsResponse;
  static deserializeBinaryFromReader(message: ListAppointmentsResponse, reader: jspb.BinaryReader): ListAppointmentsResponse;
}

export namespace ListAppointmentsResponse {
  export type AsObject = {
    appointmentsList: Array<heyrenee_v1_messages_appointment_pb.Appointment.AsObject>,
  }
}

export interface AppointmentTypeMap {
  APPOINTMENT_TYPE_UNSPECIFIED: 0;
  APPOINTMENT_UPCOMING: 1;
}

export const AppointmentType: AppointmentTypeMap;

