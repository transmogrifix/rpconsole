import {Button, ButtonGroup, createStandaloneToast, Flex, FormField, Input, PasswordInput} from '@redpanda-data/ui';
import {PageComponent, PageInitHelper} from '../../Page';
import {observer} from 'mobx-react';
import {appGlobal} from '../../../../state/appGlobal';
import {pipelinesApi, rpcnSecretManagerApi} from '../../../../state/backendApi';
import PageContent from '../../../misc/PageContent';
import {action, computed, makeObservable, observable} from 'mobx';
import {DefaultSkeleton} from '../../../../utils/tsxUtils';
import {formatPipelineError} from '../errors';
import {CreateSecretRequest, Scope} from '../../../../protogen/redpanda/api/dataplane/v1alpha2/secret_pb';
import {base64ToUInt8Array, encodeBase64} from '../../../../utils/utils';

const {ToastContainer, toast} = createStandaloneToast();

const returnSecretTab = '/connect-clusters?defaultTab=redpanda-connect-secret'


@observer
class RpConnectSecretCreate extends PageComponent {

    @observable id = '';
    @observable secret = '';
    @observable isCreating = false;

    constructor(p: any) {
        super(p);
        makeObservable(this, undefined, {autoBind: true});
    }

    initPage(p: PageInitHelper) {
        p.title = 'Create Secret';
        p.addBreadcrumb('Redpanda Connect Secret Manager', '/rp-connect/secrets/create');
        p.addBreadcrumb('Create Secret', '');

        this.refreshData(true);
        appGlobal.onRefresh = () => this.refreshData(true);
    }

    refreshData(force: boolean) {
        rpcnSecretManagerApi.refreshSecrets(force);
    }

    cancel() {
        this.secret = '';
        this.id = '';
        appGlobal.history.push(returnSecretTab);
    }

    async createSecret() {
        this.isCreating = true;

        rpcnSecretManagerApi.create(new CreateSecretRequest({
            id: this.id,
            secretData: base64ToUInt8Array(encodeBase64(this.secret)),
            scopes: [Scope.REDPANDA_CONNECT]
        }))
            .then(async () => {
                toast({
                    status: 'success', duration: 4000, isClosable: false,
                    title: 'Secret created'
                });
                await pipelinesApi.refreshPipelines(true);
                appGlobal.history.push(returnSecretTab);
            })
            .catch(err => {
                toast({
                    status: 'error', duration: null, isClosable: true,
                    title: 'Failed to create secret',
                    description: formatPipelineError(err),
                })
            })
            .finally(() => {
                this.isCreating = false;
            });
    }

    @computed
    get isNameValid() {
        if ((rpcnSecretManagerApi.secrets || []).any(x => x.id === this.id)) {
            return 'Secret name is already in use';
        }
        if (this.id === '') {
            return '';
        }
        if (!(/^[A-Z][A-Z0-9_]*$/.test(this.id))) {
            return 'The name you entered is invalid. It must start with an uppercase letter (A–Z) and can only contain uppercase letters (A–Z), digits (0–9), and underscores (_).';
        }
        if (this.id.length > 255) {
            return 'The secret name must be fewer than 255 characters.';
        }
        return '';
    }

    render() {
        if (!rpcnSecretManagerApi.secrets) return DefaultSkeleton;

        const isIdEmpty = this.id.trim().length === 0;
        const isSecretEmpty = this.secret.trim().length === 0;

        return (
            <PageContent>
                <ToastContainer/>
                <Flex flexDirection="column" gap={5}>
                    <FormField label="Secret name" isInvalid={Boolean(this.isNameValid)} errorText={this.isNameValid}>
                        <Flex alignItems="center" gap="2">
                            <Input
                                placeholder="Enter a secret name..."
                                data-testid="secretId"
                                pattern="^[A-Z][A-Z0-9_]*$"
                                min={1}
                                max={255}
                                isRequired
                                value={this.id}
                                onChange={x => this.id = x.target.value.toUpperCase()}
                                width={500}
                                disabled={this.isCreating}
                            />
                        </Flex>
                    </FormField>

                    <FormField label="Secret value">
                        <Flex alignItems="center" width={500}>
                            <PasswordInput
                                placeholder="Enter a secret value..."
                                data-testid="secretValue"
                                isRequired
                                value={this.secret}
                                onChange={x => this.secret = x.target.value}
                                width={500}
                                type="password"
                                isDisabled={this.isCreating}
                            />
                        </Flex>
                    </FormField>

                    <ButtonGroup>
                        <Button isLoading={this.isCreating} isDisabled={isIdEmpty || isSecretEmpty || Boolean(this.isNameValid)} onClick={action(() => this.createSecret())}>
                            Create Secret
                        </Button>
                        <Button variant="link" disabled={this.isCreating} onClick={action(() => this.cancel())}>
                            Cancel
                        </Button>
                    </ButtonGroup>
                </Flex>
            </PageContent>
        );

    }
}

export default RpConnectSecretCreate;
