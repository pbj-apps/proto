// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/concierge.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";

export class Concierge extends jspb.Message {
  getConciergeId(): string;
  setConciergeId(value: string): void;

  getFirstName(): string;
  setFirstName(value: string): void;

  getLastName(): string;
  setLastName(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getConciergeStatus(): ConciergeStatusMap[keyof ConciergeStatusMap];
  setConciergeStatus(value: ConciergeStatusMap[keyof ConciergeStatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Concierge.AsObject;
  static toObject(includeInstance: boolean, msg: Concierge): Concierge.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Concierge, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Concierge;
  static deserializeBinaryFromReader(message: Concierge, reader: jspb.BinaryReader): Concierge;
}

export namespace Concierge {
  export type AsObject = {
    conciergeId: string,
    firstName: string,
    lastName: string,
    email: string,
    conciergeStatus: ConciergeStatusMap[keyof ConciergeStatusMap],
  }
}

export interface ConciergeStatusMap {
  CONCIERGE_STATUS_UNSPECIFIED: 0;
  CONCIERGE_STATUS_ACTIVE: 1;
  CONCIERGE_STATUS_INACTIVE: 2;
}

export const ConciergeStatus: ConciergeStatusMap;

