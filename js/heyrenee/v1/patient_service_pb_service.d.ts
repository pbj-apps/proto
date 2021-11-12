// package: heyrenee.v1
// file: heyrenee/v1/patient_service.proto

import * as heyrenee_v1_patient_service_pb from "../../heyrenee/v1/patient_service_pb";
import * as heyrenee_v1_messages_patient_caregiver_pb from "../../heyrenee/v1/messages/patient_caregiver_pb";
import * as heyrenee_v1_messages_patient_health_questionnaire_pb from "../../heyrenee/v1/messages/patient_health_questionnaire_pb";
import * as heyrenee_v1_messages_patient_satisfaction_questionnaire_pb from "../../heyrenee/v1/messages/patient_satisfaction_questionnaire_pb";
import * as heyrenee_v1_messages_patient_sdoh_questionnaire_pb from "../../heyrenee/v1/messages/patient_sdoh_questionnaire_pb";
import * as heyrenee_v1_messages_patient_assessment_pb from "../../heyrenee/v1/messages/patient_assessment_pb";
import {grpc} from "@improbable-eng/grpc-web";

type PatientServiceCreatePatientCaregiver = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.CreatePatientCaregiverRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver;
};

type PatientServiceUpdatePatientCaregiver = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.UpdatePatientCaregiverRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver;
};

type PatientServiceListPatientCaregivers = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.ListPatientCaregiversRequest;
  readonly responseType: typeof heyrenee_v1_patient_service_pb.ListPatientCaregiversResponse;
};

type PatientServiceCreatePatientAssessment = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.CreatePatientAssessmentRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_assessment_pb.PatientAssessment;
};

type PatientServiceListPatientAssessments = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.ListPatientAssessmentsRequest;
  readonly responseType: typeof heyrenee_v1_patient_service_pb.ListPatientAssessmentsResponse;
};

type PatientServiceCreatePatientSatisfactionQuestionnaire = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.CreatePatientSatisfactionQuestionnaireRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire;
};

type PatientServiceListPatientSatisfactionQuestionnaires = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.ListPatientSatisfactionQuestionnairesRequest;
  readonly responseType: typeof heyrenee_v1_patient_service_pb.ListPatientSatisfactionQuestionnairesResponse;
};

type PatientServiceCreatePatientHealthQuestionnaire = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.CreatePatientHealthQuestionnaireRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire;
};

type PatientServiceListPatientHealthQuestionnaires = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.ListPatientHealthQuestionnairesRequest;
  readonly responseType: typeof heyrenee_v1_patient_service_pb.ListPatientHealthQuestionnairesResponse;
};

type PatientServiceCreatePatientSdohQuestionnaire = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.CreatePatientSdohQuestionnaireRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire;
};

type PatientServiceListPatientSdohQuestionnaires = {
  readonly methodName: string;
  readonly service: typeof PatientService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_patient_service_pb.ListPatientSdohQuestionnairesRequest;
  readonly responseType: typeof heyrenee_v1_patient_service_pb.ListPatientSdohQuestionnairesResponse;
};

export class PatientService {
  static readonly serviceName: string;
  static readonly CreatePatientCaregiver: PatientServiceCreatePatientCaregiver;
  static readonly UpdatePatientCaregiver: PatientServiceUpdatePatientCaregiver;
  static readonly ListPatientCaregivers: PatientServiceListPatientCaregivers;
  static readonly CreatePatientAssessment: PatientServiceCreatePatientAssessment;
  static readonly ListPatientAssessments: PatientServiceListPatientAssessments;
  static readonly CreatePatientSatisfactionQuestionnaire: PatientServiceCreatePatientSatisfactionQuestionnaire;
  static readonly ListPatientSatisfactionQuestionnaires: PatientServiceListPatientSatisfactionQuestionnaires;
  static readonly CreatePatientHealthQuestionnaire: PatientServiceCreatePatientHealthQuestionnaire;
  static readonly ListPatientHealthQuestionnaires: PatientServiceListPatientHealthQuestionnaires;
  static readonly CreatePatientSdohQuestionnaire: PatientServiceCreatePatientSdohQuestionnaire;
  static readonly ListPatientSdohQuestionnaires: PatientServiceListPatientSdohQuestionnaires;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class PatientServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createPatientCaregiver(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientCaregiverRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver|null) => void
  ): UnaryResponse;
  createPatientCaregiver(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientCaregiverRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver|null) => void
  ): UnaryResponse;
  updatePatientCaregiver(
    requestMessage: heyrenee_v1_patient_service_pb.UpdatePatientCaregiverRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver|null) => void
  ): UnaryResponse;
  updatePatientCaregiver(
    requestMessage: heyrenee_v1_patient_service_pb.UpdatePatientCaregiverRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver|null) => void
  ): UnaryResponse;
  listPatientCaregivers(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientCaregiversRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientCaregiversResponse|null) => void
  ): UnaryResponse;
  listPatientCaregivers(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientCaregiversRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientCaregiversResponse|null) => void
  ): UnaryResponse;
  createPatientAssessment(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientAssessmentRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_assessment_pb.PatientAssessment|null) => void
  ): UnaryResponse;
  createPatientAssessment(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientAssessmentRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_assessment_pb.PatientAssessment|null) => void
  ): UnaryResponse;
  listPatientAssessments(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientAssessmentsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientAssessmentsResponse|null) => void
  ): UnaryResponse;
  listPatientAssessments(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientAssessmentsRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientAssessmentsResponse|null) => void
  ): UnaryResponse;
  createPatientSatisfactionQuestionnaire(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientSatisfactionQuestionnaireRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire|null) => void
  ): UnaryResponse;
  createPatientSatisfactionQuestionnaire(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientSatisfactionQuestionnaireRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire|null) => void
  ): UnaryResponse;
  listPatientSatisfactionQuestionnaires(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientSatisfactionQuestionnairesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientSatisfactionQuestionnairesResponse|null) => void
  ): UnaryResponse;
  listPatientSatisfactionQuestionnaires(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientSatisfactionQuestionnairesRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientSatisfactionQuestionnairesResponse|null) => void
  ): UnaryResponse;
  createPatientHealthQuestionnaire(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientHealthQuestionnaireRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire|null) => void
  ): UnaryResponse;
  createPatientHealthQuestionnaire(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientHealthQuestionnaireRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire|null) => void
  ): UnaryResponse;
  listPatientHealthQuestionnaires(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientHealthQuestionnairesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientHealthQuestionnairesResponse|null) => void
  ): UnaryResponse;
  listPatientHealthQuestionnaires(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientHealthQuestionnairesRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientHealthQuestionnairesResponse|null) => void
  ): UnaryResponse;
  createPatientSdohQuestionnaire(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientSdohQuestionnaireRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire|null) => void
  ): UnaryResponse;
  createPatientSdohQuestionnaire(
    requestMessage: heyrenee_v1_patient_service_pb.CreatePatientSdohQuestionnaireRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire|null) => void
  ): UnaryResponse;
  listPatientSdohQuestionnaires(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientSdohQuestionnairesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientSdohQuestionnairesResponse|null) => void
  ): UnaryResponse;
  listPatientSdohQuestionnaires(
    requestMessage: heyrenee_v1_patient_service_pb.ListPatientSdohQuestionnairesRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_patient_service_pb.ListPatientSdohQuestionnairesResponse|null) => void
  ): UnaryResponse;
}

