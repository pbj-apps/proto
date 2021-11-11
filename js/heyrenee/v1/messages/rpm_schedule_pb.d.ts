// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/rpm_schedule.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as heyrenee_v1_messages_provider_pb from "../../../heyrenee/v1/messages/provider_pb";

export class RpmSchedule extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getRpmScheduleId(): string;
  setRpmScheduleId(value: string): void;

  getRpmScheduleType(): RpmScheduleTypeMap[keyof RpmScheduleTypeMap];
  setRpmScheduleType(value: RpmScheduleTypeMap[keyof RpmScheduleTypeMap]): void;

  hasFirstMeasurementRegimenStart(): boolean;
  clearFirstMeasurementRegimenStart(): void;
  getFirstMeasurementRegimenStart(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFirstMeasurementRegimenStart(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasMeasurementRegimenDuration(): boolean;
  clearMeasurementRegimenDuration(): void;
  getMeasurementRegimenDuration(): google_protobuf_duration_pb.Duration | undefined;
  setMeasurementRegimenDuration(value?: google_protobuf_duration_pb.Duration): void;

  getMeasurementsPerRegimen(): number;
  setMeasurementsPerRegimen(value: number): void;

  clearMeasurementDurationsFromRegimenStartList(): void;
  getMeasurementDurationsFromRegimenStartList(): Array<google_protobuf_duration_pb.Duration>;
  setMeasurementDurationsFromRegimenStartList(value: Array<google_protobuf_duration_pb.Duration>): void;
  addMeasurementDurationsFromRegimenStart(value?: google_protobuf_duration_pb.Duration, index?: number): google_protobuf_duration_pb.Duration;

  hasProviderId(): boolean;
  clearProviderId(): void;
  getProviderId(): string;
  setProviderId(value: string): void;

  hasProviderMessage(): boolean;
  clearProviderMessage(): void;
  getProviderMessage(): heyrenee_v1_messages_provider_pb.Provider | undefined;
  setProviderMessage(value?: heyrenee_v1_messages_provider_pb.Provider): void;

  getRpmScheduleStatus(): RpmScheduleStatusMap[keyof RpmScheduleStatusMap];
  setRpmScheduleStatus(value: RpmScheduleStatusMap[keyof RpmScheduleStatusMap]): void;

  getProviderCase(): RpmSchedule.ProviderCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RpmSchedule.AsObject;
  static toObject(includeInstance: boolean, msg: RpmSchedule): RpmSchedule.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RpmSchedule, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RpmSchedule;
  static deserializeBinaryFromReader(message: RpmSchedule, reader: jspb.BinaryReader): RpmSchedule;
}

export namespace RpmSchedule {
  export type AsObject = {
    patientId: string,
    rpmScheduleId: string,
    rpmScheduleType: RpmScheduleTypeMap[keyof RpmScheduleTypeMap],
    firstMeasurementRegimenStart?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    measurementRegimenDuration?: google_protobuf_duration_pb.Duration.AsObject,
    measurementsPerRegimen: number,
    measurementDurationsFromRegimenStartList: Array<google_protobuf_duration_pb.Duration.AsObject>,
    providerId: string,
    providerMessage?: heyrenee_v1_messages_provider_pb.Provider.AsObject,
    rpmScheduleStatus: RpmScheduleStatusMap[keyof RpmScheduleStatusMap],
  }

  export enum ProviderCase {
    PROVIDER_NOT_SET = 0,
    PROVIDER_ID = 8,
    PROVIDER_MESSAGE = 9,
  }
}

export interface RpmScheduleTypeMap {
  RPM_SCHEDULE_TYPE_UNSPECIFIED: 0;
  RPM_SCHEDULE_TYPE_HEART_RATE: 1;
  RPM_SCHEDULE_TYPE_BLOOD_PRESSURE: 2;
  RPM_SCHEDULE_TYPE_PULSE: 3;
  RPM_SCHEDULE_TYPE_SP_O2: 4;
  RPM_SCHEDULE_TYPE_WEIGHT: 5;
  RPM_SCHEDULE_TYPE_GLUCOSE: 6;
}

export const RpmScheduleType: RpmScheduleTypeMap;

export interface RpmScheduleStatusMap {
  RPM_SCHEDULE_STATUS_UNSPECIFIED: 0;
  RPM_SCHEDULE_ACTIVE: 1;
  RPM_SCHEDULE_INACTIVE: 2;
}

export const RpmScheduleStatus: RpmScheduleStatusMap;

