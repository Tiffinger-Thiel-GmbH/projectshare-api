import { get, post } from './common';

export interface ProjectDTO {
  id: string;
  name: string;
}

export interface DocumentDTO {
  id: string;
  location: string;
  name: string;
}

export const baseUrl = 'http://localhost:7777';

export const getProjects = async (): Promise<ProjectDTO[]> => {
  const projects = await get(baseUrl + '/project').then(response => response.data);
  // .catch(() => [
  //   { id: '0975ea14-56a0-4d40-940c-c9e94aa6b359', name: 'ProjectX' },
  //   { id: '22ab3304-5fd8-45ca-b7bd-0ae9836f4e28', name: 'ProjectY' }
  // ]);
  return projects;
};

export const getDocumentsByProjectId = async (projectId: string): Promise<DocumentDTO[]> => {
  const documents = await get(baseUrl + '/project/' + projectId + '/document').then(response => response.data);
  // .catch(() => [{ id: '8eff87f7-6528-48fc-b4ac-e5e22fd195d9', bucket: '22ab3304-5fd8-45ca-b7bd-0ae9836f4e28', name: 'Plan1.txt' }]);

  return documents;
};

/**
 *
 * @param documentId the id of the document you want to download
 * @param locationId the id of the location where the desired document is stored (projectId)
 */
export const getDocumentByDocumentId = async (documentId: string, locationId: string): Promise<BlobPart> => {
  const document = await get(baseUrl + '/document/' + locationId + '/' + documentId, { responseType: 'blob' }).then(
    response => response.data
  );
  return document;
};

export const uploadDocument = async (locationId: string, file: FormData): Promise<void> => {
  const result = await post(baseUrl + '/document/' + locationId, file).then(response => response.data);

  return result;
};

export const createProject = async (name: string): Promise<void> => {
  const result = await post(baseUrl + '/project', { name }).then(response => response.data);

  return result;
};
