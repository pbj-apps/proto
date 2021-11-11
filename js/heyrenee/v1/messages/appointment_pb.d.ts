// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/appointment.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_provider_pb from "../../../heyrenee/v1/messages/provider_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Appointment extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  hasProviderId(): boolean;
  clearProviderId(): void;
  getProviderId(): string;
  setProviderId(value: string): void;

  hasProviderMessage(): boolean;
  clearProviderMessage(): void;
  getProviderMessage(): heyrenee_v1_messages_provider_pb.Provider | undefined;
  setProviderMessage(value?: heyrenee_v1_messages_provider_pb.Provider): void;

  getAppointmentId(): string;
  setAppointmentId(value: string): void;

  hasDate(): boolean;
  clearDate(): void;
  getDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  clearAddressLinesList(): void;
  getAddressLinesList(): Array<string>;
  setAddressLinesList(value: Array<string>): void;
  addAddressLines(value: string, index?: number): string;

  getCity(): string;
  setCity(value: string): void;

  getState(): string;
  setState(value: string): void;

  getZip(): string;
  setZip(value: string): void;

  getInstructions(): string;
  setInstructions(value: string): void;

  getProviderCase(): Appointment.ProviderCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Appointment.AsObject;
  static toObject(includeInstance: boolean, msg: Appointment): Appointment.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Appointment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Appointment;
  static deserializeBinaryFromReader(message: Appointment, reader: jspb.BinaryReader): Appointment;
}

export namespace Appointment {
  export type AsObject = {
    patientId: string,
    providerId: string,
    providerMessage?: heyrenee_v1_messages_provider_pb.Provider.AsObject,
    appointmentId: string,
    date?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    addressLinesList: Array<string>,
    city: string,
    state: string,
    zip: string,
    instructions: string,
  }

  export enum ProviderCase {
    PROVIDER_NOT_SET = 0,
    PROVIDER_ID = 2,
    PROVIDER_MESSAGE = 3,
  }
}

